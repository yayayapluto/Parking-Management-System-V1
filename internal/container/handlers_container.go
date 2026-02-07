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

	// Customer & Vehicle Handlers
	c.CustomerHandler = handlers.NewCustomerHandler(c.CustomerService)
	c.VehicleHandler = handlers.NewVehicleHandler(c.VehicleService)

	// Auth & User Handlers
	c.UserHandler = handlers.NewUserHandler(c.UserService)
	c.RoleHandler = handlers.NewRoleHandler(c.RoleService)
	c.AuthHandler = handlers.NewAuthHandler(c.AuthService, cfg)

	// Transaction Handlers
	c.ParkingTransactionHandler = handlers.NewParkingTransactionHandler(c.ParkingTransactionService, c.IDObfuscator)

	// Financials Handlers
	c.RefundHandler = handlers.NewRefundHandler(c.RefundService)
	c.LostTicketFeeHandler = handlers.NewLostTicketFeeHandler(c.LostTicketFeeService)

	// Operations Handlers
	c.ManualCorrectionHandler = handlers.NewManualCorrectionHandler(c.ManualCorrectionService)
	c.TransactionEventHandler = handlers.NewTransactionEventHandler(c.TransactionEventService)

	// Reporting Handlers
	c.ShiftReportHandler = handlers.NewShiftReportHandler(c.ShiftReportService)
	c.DailySettlementHandler = handlers.NewDailySettlementHandler(c.DailySettlementService)

	// Logging Handlers
	c.OCRLogHandler = handlers.NewOCRLogHandler(c.OCRLogService)
	c.UserActivityLogHandler = handlers.NewUserActivityLogHandler(c.UserActivityLogService)
	c.SystemLogHandler = handlers.NewSystemLogHandler(c.SystemLogService)
	c.ZoneOccupancyLogHandler = handlers.NewZoneOccupancyLogHandler(c.ZoneOccupancyLogService)

	// Others Handlers
	c.PermissionHandler = handlers.NewPermissionHandler(c.PermissionService)
	c.ReportCacheHandler = handlers.NewReportCacheHandler(c.ReportCacheService)
}
