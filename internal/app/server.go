package app

import (
	"errors"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"os/signal"
	"parking-management-system-v1/internal/container"
	"parking-management-system-v1/internal/dto/responses"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"parking-management-system-v1/internal/routes"
	"parking-management-system-v1/pkg/config"
)

type Server struct {
	Fiber     *fiber.App
	Container *container.Container
	Config    *config.Config
}

func NewServer(cfg *config.Config, container *container.Container) *Server {
	app := fiber.New(fiber.Config{
		AppName:           "Parking Management System v1",
		ErrorHandler:      globalErrorHandler, // Fungsi error handler pindahin ke bawah atau pkg
		EnablePrintRoutes: true,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		BodyLimit:         10 * 1024 * 1024, // 10mb
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
	}))
	app.Get("/metrics", monitor.New())

	// Init Routes
	routes.SetupRoutes(app, container)

	return &Server{
		Fiber:     app,
		Container: container,
		Config:    cfg,
	}
}

func (s *Server) Run() {
	// Jalankan server di goroutine
	go func() {
		addr := ":" + s.Config.AppConfig.AppPort // Pastikan ada AppPort di config
		if err := s.Fiber.Listen(addr); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Printf("Server is running on port %s", s.Config.AppConfig.AppPort)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	if err := s.Fiber.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

func globalErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code default (Internal Server Error)
	code := fiber.StatusInternalServerError

	// Cek jika error-nya adalah Fiber Error (misal 404 Not Found atau 400 Bad Request)
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Response seragam pakai BaseResponse
	return ctx.Status(code).JSON(responses.BaseResponse{
		Success: false,
		Message: "An error occurred",
		Data:    nil,
		Errors:  err.Error(),
	})
}
