package app

import (
	"errors"
	"github.com/go-playground/validator/v10"
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

	customValidator "parking-management-system-v1/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "parking-management-system-v1/docs"
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

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Init Routes
	routes.SetupRoutes(app, container, cfg)

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
	code := fiber.StatusInternalServerError
	message := "An error occurred"
	var errorData interface{} = err.Error()

	// 1. Tangkap Validator Errors
	var valErrors validator.ValidationErrors
	if errors.As(err, &valErrors) {
		code = fiber.StatusBadRequest
		message = "Validation failed" // Tambahin biar jelas di response
		errFields := make(map[string]string)

		for _, e := range valErrors {
			// PANGGIL PAKE ALIAS: customValidator.Trans
			errFields[e.Field()] = e.Translate(customValidator.Trans)
		}
		errorData = errFields
	}

	// 2. Cek kalau ini Error dari Fiber
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		if code == fiber.StatusUnprocessableEntity {
			code = fiber.StatusBadRequest
		}
	}

	return ctx.Status(code).JSON(responses.BaseResponse{
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  errorData,
	})
}
