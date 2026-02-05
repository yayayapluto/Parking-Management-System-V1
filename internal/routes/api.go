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

	// VehicleType Routes
	vt := v1.Group("/vehicle-types")
	vt.Get("/info", DiscoveryHandler(fb))
	vt.Get("/", c.VehicleTypeHandler.GetAll)
	vt.Post("/", c.VehicleTypeHandler.Create)
	vt.Get("/:id", c.VehicleTypeHandler.GetByID)
	vt.Put("/:id", c.VehicleTypeHandler.Update)
	vt.Delete("/:id", c.VehicleTypeHandler.Delete)

	// ZoneType Routes
	zt := v1.Group("/zone-types")
	zt.Get("/info", DiscoveryHandler(fb))
	zt.Get("/", c.ZoneTypeHandler.GetAll)
	zt.Post("/", c.ZoneTypeHandler.Create)
	zt.Get("/:id", c.ZoneTypeHandler.GetByID)
	zt.Put("/:id", c.ZoneTypeHandler.Update)
	zt.Delete("/:id", c.ZoneTypeHandler.Delete)
}
