package app

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/repositories"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/helpers"
)

type Container struct {
	VehicleTypeRepo repositories.VehicleTypeRepository
	VehicleRepo     repositories.VehicleRepository

	PermissionRepo repositories.PermissionRepository
	RoleRepo       repositories.RoleRepository
	UserRepo       repositories.UserRepository

	ZoneTypeRepo repositories.ZoneTypeRepository
	ZoneRepo     repositories.ZoneRepository
	ZoneRateRepo repositories.ZoneRateRepository

	PaymentMethodRepo repositories.PaymentMethodRepository

	HolidayRepo repositories.HolidayRepository

	CustomerRegistSourceRepo repositories.CustomerRegistSourceRepository
	CustomerRepo             repositories.CustomerRepository

	ParkingTransactionRepo repositories.ParkingTransactionRepository

	TransactionZoneRepo    repositories.TransactionZoneRepository
	TransactionOCRDataRepo repositories.TransactionOCRDataRepository
	TransactionEventRepo   repositories.TransactionEventRepository

	PaymentRepo       repositories.PaymentRepository
	RefundRepo        repositories.RefundRepository
	LostTicketFeeRepo repositories.LostTicketFeeRepository

	OCRLogRepo          repositories.OCRLogRepository
	UserActivityLogRepo repositories.UserActivityLogRepository
	SystemLogRepo       repositories.SystemLogRepository

	ZoneOccupancyLogRepo repositories.ZoneOccupancyLogRepository
	ManualCorrectionRepo repositories.ManualCorrectionRepository
	ShiftReportRepo      repositories.ShiftReportRepository

	DailySettlementRepo repositories.DailySettlementRepository
	ReportCacheRepo     repositories.ReportCacheRepository

	IDObfuscator helpers.IDObfuscator
}

func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	obfuscator, _ := helpers.NewHashIDManager(cfg.HashConfig.Salt, cfg.HashConfig.MinLength)
	return &Container{
		IDObfuscator: obfuscator,

		// Group: Vehicle
		VehicleTypeRepo: repositories.NewVehicleTypeRepository(db),
		VehicleRepo:     repositories.NewVehicleRepository(db),

		// Group: Auth & User
		PermissionRepo: repositories.NewPermissionRepository(db),
		RoleRepo:       repositories.NewRoleRepository(db),
		UserRepo:       repositories.NewUserRepository(db),

		// Group: Zone & Rates
		ZoneTypeRepo: repositories.NewZoneTypeRepository(db),
		ZoneRepo:     repositories.NewZoneRepository(db),
		ZoneRateRepo: repositories.NewZoneRateRepository(db),

		// Group: Payment Config
		PaymentMethodRepo: repositories.NewPaymentMethodRepository(db),

		// Group: General
		HolidayRepo: repositories.NewHolidayRepository(db),

		// Group: Customer
		CustomerRegistSourceRepo: repositories.NewCustomerRegistSourceRepository(db),
		CustomerRepo:             repositories.NewCustomerRepository(db),

		// Group: Main Transaction
		ParkingTransactionRepo: repositories.NewParkingTransactionRepository(db),

		// Group: Transaction Details
		TransactionZoneRepo:    repositories.NewTransactionZoneRepository(db),
		TransactionOCRDataRepo: repositories.NewTransactionOCRDataRepository(db),
		TransactionEventRepo:   repositories.NewTransactionEventRepository(db),

		// Group: Financials
		PaymentRepo:       repositories.NewPaymentRepository(db),
		RefundRepo:        repositories.NewRefundRepository(db),
		LostTicketFeeRepo: repositories.NewLostTicketFeeRepository(db),

		// Group: Logs
		OCRLogRepo:          repositories.NewOCRLogRepository(db),
		UserActivityLogRepo: repositories.NewUserActivityLogRepository(db),
		SystemLogRepo:       repositories.NewSystemLogRepository(db),

		// Group: Operations
		ZoneOccupancyLogRepo: repositories.NewZoneOccupancyLogRepository(db),
		ManualCorrectionRepo: repositories.NewManualCorrectionRepository(db),
		ShiftReportRepo:      repositories.NewShiftReportRepository(db),

		// Group: Reports
		DailySettlementRepo: repositories.NewDailySettlementRepository(db),
		ReportCacheRepo:     repositories.NewReportCacheRepository(db),
	}
}
