package container

import (
	"parking-management-system-v1/internal/handlers"
	"parking-management-system-v1/pkg/config"
)

func initHandlers(c *Container, cfg *config.Config) {
	c.VehicleTypeHandler = handlers.NewVehicleTypeHandler(c.VehicleTypeService)
	c.ZoneHandler = handlers.NewZoneHandler(c.ZoneService)
	c.ZoneRateHandler = handlers.NewZoneRateHandler(c.ZoneRateService)
	c.ZoneTypeHandler = handlers.NewZoneTypeHandler(c.ZoneTypeService)
	c.CustomerRegistSourceHandler = handlers.NewCustomerRegistSourceHandler(c.CustomerRegistSourceService)
	c.HolidayHandler = handlers.NewHolidayHandler(c.HolidayService)
	c.PaymentMethodHandler = handlers.NewPaymentMethodHandler(c.PaymentMethodService)
	
	// Auth & User Handlers
	c.UserHandler = handlers.NewUserHandler(c.UserService)
	c.RoleHandler = handlers.NewRoleHandler(c.RoleService)
	c.AuthHandler = handlers.NewAuthHandler(c.AuthService, cfg)
	
	// Transaction Handlers
	c.ParkingTransactionHandler = handlers.NewParkingTransactionHandler(c.ParkingTransactionService, c.IDObfuscator)
}
