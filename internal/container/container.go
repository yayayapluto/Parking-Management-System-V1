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
	PermissionRepo repos.PermissionRepository
	RoleRepo       repos.RoleRepository
	UserRepo       repos.UserRepository

	// Group: Vehicle
	VehicleRepo        repos.VehicleRepository
	VehicleTypeRepo    repos.VehicleTypeRepository
	VehicleTypeService services.VehicleTypeService
	VehicleTypeHandler *handlers.VehicleTypeHandler

	// Group: Zone & Rates
	ZoneRepo        repos.ZoneRepository
	ZoneRateRepo    repos.ZoneRateRepository
	ZoneTypeRepo    repos.ZoneTypeRepository
	ZoneTypeService services.ZoneTypeService

	// Group: Customer
	CustomerRepo                repos.CustomerRepository
	CustomerRegistSourceRepo    repos.CustomerRegistSourceRepository
	CustomerRegistSourceService services.CustomerRegistSourceService

	// Group: Payment & General
	HolidayRepo          repos.HolidayRepository
	HolidayService       services.HolidayService
	PaymentMethodRepo    repos.PaymentMethodRepository
	PaymendMethodService services.PaymentMethodService

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
	obfuscator, err := helpers.NewHashIDManager(cfg.HashConfig.Salt, cfg.HashConfig.MinLength)
	if err != nil {
		panic("Failed to init HashID: " + err.Error())
	}

	validate := validator.NewValidator()

	// Initialize Repositories
	permissionRepo := repos.NewPermissionRepository(db)
	roleRepo := repos.NewRoleRepository(db)
	userRepo := repos.NewUserRepository(db)

	vehicleRepo := repos.NewVehicleRepository(db)
	vehicleTypeRepo := repos.NewVehicleTypeRepository(db)

	zoneRepo := repos.NewZoneRepository(db)
	zoneRateRepo := repos.NewZoneRateRepository(db)
	zoneTypeRepo := repos.NewZoneTypeRepository(db)

	customerRepo := repos.NewCustomerRepository(db)
	customerRegistSourceRepo := repos.NewCustomerRegistSourceRepository(db)

	holidayRepo := repos.NewHolidayRepository(db)
	paymentMethodRepo := repos.NewPaymentMethodRepository(db)

	// Initialize Services
	vehicleTypeService := services.NewVehicleTypeService(vehicleTypeRepo, obfuscator, validate)
	zoneTypeService := services.NewZoneTypeService(zoneTypeRepo, obfuscator, validate)
	holidayService := services.NewHolidayService(holidayRepo, obfuscator, validate)
	customerRegistSourceService := services.NewCustomerRegistSourceService(customerRegistSourceRepo, obfuscator, validate)
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepo, obfuscator, validate)

	// Initialize Handlers
	vehicleTypeHandler := handlers.NewVehicleTypeHandler(vehicleTypeService)

	return &Container{
		IDObfuscator: obfuscator,

		// Group: Auth & User
		PermissionRepo: permissionRepo,
		RoleRepo:       roleRepo,
		UserRepo:       userRepo,

		// Group: Vehicle
		VehicleRepo:        vehicleRepo,
		VehicleTypeRepo:    vehicleTypeRepo,
		VehicleTypeService: vehicleTypeService,
		VehicleTypeHandler: vehicleTypeHandler,

		// Group: Zone & Rates
		ZoneRepo:        zoneRepo,
		ZoneRateRepo:    zoneRateRepo,
		ZoneTypeRepo:    zoneTypeRepo,
		ZoneTypeService: zoneTypeService,

		// Group: Customer
		CustomerRepo:                customerRepo,
		CustomerRegistSourceRepo:    customerRegistSourceRepo,
		CustomerRegistSourceService: customerRegistSourceService,

		// Group: Payment & General
		HolidayRepo:          holidayRepo,
		HolidayService:       holidayService,
		PaymentMethodRepo:    paymentMethodRepo,
		PaymendMethodService: paymentMethodService,

		// Group: Main Transaction
		ParkingTransactionRepo: repos.NewParkingTransactionRepository(db),
		TransactionZoneRepo:    repos.NewTransactionZoneRepository(db),
		TransactionOCRDataRepo: repos.NewTransactionOCRDataRepository(db),
		TransactionEventRepo:   repos.NewTransactionEventRepository(db),

		// Group: Financials
		PaymentRepo:       repos.NewPaymentRepository(db),
		RefundRepo:        repos.NewRefundRepository(db),
		LostTicketFeeRepo: repos.NewLostTicketFeeRepository(db),

		// Group: Operations & Logs
		ManualCorrectionRepo: repos.NewManualCorrectionRepository(db),
		ShiftReportRepo:      repos.NewShiftReportRepository(db),
		ZoneOccupancyLogRepo: repos.NewZoneOccupancyLogRepository(db),
		OCRLogRepo:           repos.NewOCRLogRepository(db),
		SystemLogRepo:        repos.NewSystemLogRepository(db),
		UserActivityLogRepo:  repos.NewUserActivityLogRepository(db),

		// Group: Reports
		DailySettlementRepo: repos.NewDailySettlementRepository(db),
		ReportCacheRepo:     repos.NewReportCacheRepository(db),
	}
}
