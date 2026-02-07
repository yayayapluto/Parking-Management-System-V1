package container

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/handlers"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/helpers"
	"parking-management-system-v1/pkg/validator"
)

type Container struct {
	IDObfuscator helpers.IDObfuscator

	// Group: Auth & User
	PermissionRepo    repos.PermissionRepository
	RoleRepo          repos.RoleRepository
	UserRepo          repos.UserRepository
	UserService       services.UserService
	RoleService       services.RoleService
	AuthService       services.AuthService
	UserHandler       *handlers.UserHandler
	RoleHandler       *handlers.RoleHandler
	AuthHandler       *handlers.AuthHandler

	// Group: Vehicle
	VehicleRepo        repos.VehicleRepository
	VehicleTypeRepo    repos.VehicleTypeRepository
	VehicleTypeService services.VehicleTypeService
	VehicleTypeHandler *handlers.VehicleTypeHandler

	// Group: Zone & Rates
	ZoneRepo        repos.ZoneRepository
	ZoneRateRepo    repos.ZoneRateRepository
	ZoneTypeRepo    repos.ZoneTypeRepository
	ZoneService     services.ZoneService
	ZoneHandler     *handlers.ZoneHandler
	ZoneRateService services.ZoneRateService
	ZoneRateHandler *handlers.ZoneRateHandler
	ZoneTypeService services.ZoneTypeService
	ZoneTypeHandler *handlers.ZoneTypeHandler

	// Group: Customer
	CustomerRepo                repos.CustomerRepository
	CustomerRegistSourceRepo    repos.CustomerRegistSourceRepository
	CustomerRegistSourceService services.CustomerRegistSourceService
	CustomerRegistSourceHandler *handlers.CustomerRegistSourceHandler

	// Group: Payment & General
	HolidayRepo          repos.HolidayRepository
	HolidayService       services.HolidayService
	HolidayHandler       *handlers.HolidayHandler
	PaymentMethodRepo    repos.PaymentMethodRepository
	PaymentMethodService services.PaymentMethodService
	PaymentMethodHandler *handlers.PaymentMethodHandler

	// Group: Main Transaction & Details
	ParkingTransactionRepo repos.ParkingTransactionRepository
	TransactionZoneRepo    repos.TransactionZoneRepository
	TransactionOCRDataRepo repos.TransactionOCRDataRepository
	TransactionEventRepo   repos.TransactionEventRepository

	// Group: Financials
	PaymentRepo       repos.PaymentRepository
	RefundRepo        repos.RefundRepository
	LostTicketFeeRepo repos.LostTicketFeeRepository

	// Group: Operations & Logs
	ManualCorrectionRepo repos.ManualCorrectionRepository
	ShiftReportRepo      repos.ShiftReportRepository
	ZoneOccupancyLogRepo repos.ZoneOccupancyLogRepository
	OCRLogRepo           repos.OCRLogRepository
	SystemLogRepo        repos.SystemLogRepository
	UserActivityLogRepo  repos.UserActivityLogRepository

	// Group: Reports
	DailySettlementRepo repos.DailySettlementRepository
	ReportCacheRepo     repos.ReportCacheRepository
}

func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	c := &Container{}

	// Utils
	obfuscator, _ := helpers.NewHashIDManager(cfg.HashConfig.Salt, cfg.HashConfig.MinLength)
	validate := validator.NewValidator()
	c.IDObfuscator = obfuscator

	// Layer Initialization
	initRepositories(db, c)
	initServices(c, obfuscator, validate)
	initHandlers(c, cfg)

	return c
}
