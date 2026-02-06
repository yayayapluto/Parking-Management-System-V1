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

	cr := v1.Group("/customer-regist-sources")
	cr.Get("/info", DiscoveryHandler(fb))
	cr.Get("/", c.CustomerRegistSourceHandler.GetAll)
	cr.Post("/", c.CustomerRegistSourceHandler.Create)
	cr.Get("/:id", c.CustomerRegistSourceHandler.GetByID)
	cr.Put("/:id", c.CustomerRegistSourceHandler.Update)
	cr.Delete("/:id", c.CustomerRegistSourceHandler.Delete)

	h := v1.Group("/holidays")
	h.Get("/info", DiscoveryHandler(fb))
	h.Get("/", c.HolidayHandler.GetAll)
	h.Post("/", c.HolidayHandler.Create)
	h.Get("/:id", c.HolidayHandler.GetByID)
	h.Put("/:id", c.HolidayHandler.Update)
	h.Delete("/:id", c.HolidayHandler.Delete)

	pm := v1.Group("/payment-methods")
	pm.Get("/info", DiscoveryHandler(fb))
	pm.Get("/", c.PaymentMethodHandler.GetAll)
	pm.Post("/", c.PaymentMethodHandler.Create)
	pm.Get("/:id", c.PaymentMethodHandler.GetByID)
	pm.Put("/:id", c.PaymentMethodHandler.Update)
	pm.Delete("/:id", c.PaymentMethodHandler.Delete)
}
