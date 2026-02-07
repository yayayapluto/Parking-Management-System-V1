package routes

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/container"
	"parking-management-system-v1/pkg/config"
)

func SetupRoutes(fb *fiber.App, c *container.Container, cfg *config.Config) {
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

	// Zone Routes
	z := v1.Group("/zones")
	z.Get("/info", DiscoveryHandler(fb))
	z.Get("/", c.ZoneHandler.GetAll)
	z.Post("/", c.ZoneHandler.Create)
	z.Get("/:id", c.ZoneHandler.GetByID)
	z.Put("/:id", c.ZoneHandler.Update)
	z.Delete("/:id", c.ZoneHandler.Delete)

	// ZoneRate Routes
	zr := v1.Group("/zone-rates")
	zr.Get("/info", DiscoveryHandler(fb))
	zr.Get("/", c.ZoneRateHandler.GetAll)
	zr.Post("/", c.ZoneRateHandler.Create)
	zr.Get("/:id", c.ZoneRateHandler.GetByID)
	zr.Put("/:id", c.ZoneRateHandler.Update)
	zr.Delete("/:id", c.ZoneRateHandler.Delete)

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

	// Customer Routes
	cust := v1.Group("/customers")
	cust.Get("/info", DiscoveryHandler(fb))
	cust.Get("/", c.CustomerHandler.GetAll)
	cust.Post("/", c.CustomerHandler.Create)
	cust.Get("/:id", c.CustomerHandler.GetByID)
	cust.Put("/:id", c.CustomerHandler.Update)
	cust.Delete("/:id", c.CustomerHandler.Delete)

	// Vehicle Routes
	veh := v1.Group("/vehicles")
	veh.Get("/info", DiscoveryHandler(fb))
	veh.Get("/", c.VehicleHandler.GetAll)
	veh.Post("/", c.VehicleHandler.Create)
	veh.Get("/:id", c.VehicleHandler.GetByID)
	veh.Put("/:id", c.VehicleHandler.Update)
	veh.Delete("/:id", c.VehicleHandler.Delete)

	// Auth Routes (No JWT required)
	auth := v1.Group("/auth")
	auth.Get("/info", DiscoveryHandler(fb))
	auth.Post("/login", c.AuthHandler.Login)
	auth.Post("/login-username", c.AuthHandler.LoginWithUsername)

	// Master Routes (Users & Roles) - Protected by JWT
	master := v1.Group("/")
	//master.Use(middlewares.JWTMiddleware(cfg))

	// User Routes
	users := master.Group("/users")
	users.Get("/info", DiscoveryHandler(fb))
	users.Get("/", c.UserHandler.GetAll)
	users.Post("/", c.UserHandler.Create)
	users.Get("/:id", c.UserHandler.GetByID)
	users.Put("/:id", c.UserHandler.Update)
	users.Delete("/:id", c.UserHandler.Delete)

	// Role Routes
	roles := master.Group("/roles")
	roles.Get("/info", DiscoveryHandler(fb))
	roles.Get("/", c.RoleHandler.GetAll)
	roles.Post("/", c.RoleHandler.Create)
	roles.Get("/:id", c.RoleHandler.GetByID)
	roles.Put("/:id", c.RoleHandler.Update)
	roles.Delete("/:id", c.RoleHandler.Delete)

	// Transaction Routes
	transactions := v1.Group("/transactions")
	transactions.Get("/info", DiscoveryHandler(fb))
	transactions.Post("/entry", c.ParkingTransactionHandler.Entry)
	transactions.Post("/exit", c.ParkingTransactionHandler.Exit)

	// ============================================
	// FINANCIALS ROUTES
	// ============================================
	financials := v1.Group("/financials")

	// Refund Routes
	refunds := financials.Group("/refunds")
	refunds.Get("/info", DiscoveryHandler(fb))
	refunds.Get("/", c.RefundHandler.GetAll)
	refunds.Post("/", c.RefundHandler.Create)
	refunds.Get("/:id", c.RefundHandler.GetByID)
	refunds.Put("/:id", c.RefundHandler.Update)
	refunds.Post("/:id/approve", c.RefundHandler.Approve)
	refunds.Post("/:id/reject", c.RefundHandler.Reject)
	refunds.Delete("/:id", c.RefundHandler.Delete)

	// Lost Ticket Fee Routes
	ltf := financials.Group("/lost-ticket-fees")
	ltf.Get("/info", DiscoveryHandler(fb))
	ltf.Get("/", c.LostTicketFeeHandler.GetAll)
	ltf.Post("/", c.LostTicketFeeHandler.Create)
	ltf.Get("/:id", c.LostTicketFeeHandler.GetByID)
	ltf.Delete("/:id", c.LostTicketFeeHandler.Delete)

	// ============================================
	// OPERATIONS ROUTES
	// ============================================
	operations := v1.Group("/operations")

	// Manual Correction Routes
	mc := operations.Group("/manual-corrections")
	mc.Get("/info", DiscoveryHandler(fb))
	mc.Get("/", c.ManualCorrectionHandler.GetAll)
	mc.Post("/", c.ManualCorrectionHandler.Create)
	mc.Get("/:id", c.ManualCorrectionHandler.GetByID)
	mc.Delete("/:id", c.ManualCorrectionHandler.Delete)

	// Transaction Event Routes
	te := operations.Group("/transaction-events")
	te.Get("/info", DiscoveryHandler(fb))
	te.Get("/", c.TransactionEventHandler.GetAll)
	te.Post("/", c.TransactionEventHandler.Create)
	te.Get("/:id", c.TransactionEventHandler.GetByID)
	te.Delete("/:id", c.TransactionEventHandler.Delete)

	// ============================================
	// ADMIN ROUTES
	// ============================================
	admin := v1.Group("/admin")

	// Reports Routes
	reports := admin.Group("/reports")

	// Shift Report Routes
	sr := reports.Group("/shift-reports")
	sr.Get("/info", DiscoveryHandler(fb))
	sr.Get("/", c.ShiftReportHandler.GetAll)
	sr.Post("/", c.ShiftReportHandler.Create)
	sr.Get("/:id", c.ShiftReportHandler.GetByID)
	sr.Put("/:id", c.ShiftReportHandler.Update)
	sr.Delete("/:id", c.ShiftReportHandler.Delete)

	// Daily Settlement Routes
	ds := reports.Group("/daily-settlements")
	ds.Get("/info", DiscoveryHandler(fb))
	ds.Get("/", c.DailySettlementHandler.GetAll)
	ds.Post("/", c.DailySettlementHandler.Create)
	ds.Get("/:id", c.DailySettlementHandler.GetByID)
	ds.Post("/:id/reconcile", c.DailySettlementHandler.Reconcile)
	ds.Delete("/:id", c.DailySettlementHandler.Delete)

	// Report Cache Routes
	rc := reports.Group("/cache")
	rc.Get("/info", DiscoveryHandler(fb))
	rc.Get("/", c.ReportCacheHandler.GetAll)
	rc.Post("/", c.ReportCacheHandler.Create)
	rc.Get("/:id", c.ReportCacheHandler.GetByID)
	rc.Put("/:id", c.ReportCacheHandler.Update)
	rc.Delete("/:id", c.ReportCacheHandler.Delete)

	// Logs Routes
	logs := admin.Group("/logs")

	// OCR Log Routes
	ocr := logs.Group("/ocr-logs")
	ocr.Get("/info", DiscoveryHandler(fb))
	ocr.Get("/", c.OCRLogHandler.GetAll)
	ocr.Post("/", c.OCRLogHandler.Create)
	ocr.Get("/:id", c.OCRLogHandler.GetByID)
	ocr.Delete("/:id", c.OCRLogHandler.Delete)

	// User Activity Log Routes
	ual := logs.Group("/user-activity-logs")
	ual.Get("/info", DiscoveryHandler(fb))
	ual.Get("/", c.UserActivityLogHandler.GetAll)
	ual.Post("/", c.UserActivityLogHandler.Create)
	ual.Get("/:id", c.UserActivityLogHandler.GetByID)
	ual.Delete("/:id", c.UserActivityLogHandler.Delete)

	// System Log Routes
	sl := logs.Group("/system-logs")
	sl.Get("/info", DiscoveryHandler(fb))
	sl.Get("/", c.SystemLogHandler.GetAll)
	sl.Post("/", c.SystemLogHandler.Create)
	sl.Get("/:id", c.SystemLogHandler.GetByID)
	sl.Delete("/:id", c.SystemLogHandler.Delete)

	// Zone Occupancy Log Routes
	zol := logs.Group("/zone-occupancy-logs")
	zol.Get("/info", DiscoveryHandler(fb))
	zol.Get("/", c.ZoneOccupancyLogHandler.GetAll)
	zol.Post("/", c.ZoneOccupancyLogHandler.Create)
	zol.Get("/:id", c.ZoneOccupancyLogHandler.GetByID)
	zol.Delete("/:id", c.ZoneOccupancyLogHandler.Delete)

	// Permission Routes
	perms := admin.Group("/permissions")
	perms.Get("/info", DiscoveryHandler(fb))
	perms.Get("/", c.PermissionHandler.GetAll)
	perms.Post("/", c.PermissionHandler.Create)
	perms.Get("/:id", c.PermissionHandler.GetByID)
	perms.Put("/:id", c.PermissionHandler.Update)
	perms.Delete("/:id", c.PermissionHandler.Delete)
}
