package container

import (
	"parking-management-system-v1/internal/handlers"
)

func initHandlers(c *Container) {
	c.VehicleTypeHandler = handlers.NewVehicleTypeHandler(c.VehicleTypeService)
	c.ZoneTypeHandler = handlers.NewZoneTypeHandler(c.ZoneTypeService)
}
