package routes

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/container"
)

func SetupRoutes(fb *fiber.App, c *container.Container) {
	// Root Discovery - Panggil fungsi discovery tadi
	fb.Get("/", DiscoveryHandler(fb))

	api := fb.Group("/api")
	api.Get("/", DiscoveryHandler(fb))

	v1 := api.Group("/v1")
	v1.Get("/", DiscoveryHandler(fb))

	// Vehicle Type Routes
	vt := v1.Group("/vehicle-types")
	vt.Get("/info", DiscoveryHandler(fb)) // Discovery level vehicle-types
	vt.Get("/", c.VehicleTypeHandler.GetAll)
	vt.Post("/", c.VehicleTypeHandler.Create)
	vt.Get("/{id}", c.VehicleTypeHandler.GetByID)
}
