package container

import (
	"github.com/go-playground/validator/v10"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

func initServices(c *Container, o helpers.IDObfuscator, v *validator.Validate) {
	c.VehicleTypeService = services.NewVehicleTypeService(c.VehicleTypeRepo, o, v)
	c.ZoneTypeService = services.NewZoneTypeService(c.ZoneTypeRepo, o, v)
	c.HolidayService = services.NewHolidayService(c.HolidayRepo, o, v)
	c.CustomerRegistSourceService = services.NewCustomerRegistSourceService(c.CustomerRegistSourceRepo, o, v)
	c.PaymentMethodService = services.NewPaymentMethodService(c.PaymentMethodRepo, o, v)
}
