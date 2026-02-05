package container

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/repos"
)

func initRepositories(db *gorm.DB, c *Container) {
	// Auth & User
	c.PermissionRepo = repos.NewPermissionRepository(db)
	c.RoleRepo = repos.NewRoleRepository(db)
	c.UserRepo = repos.NewUserRepository(db)

	// Vehicle & Zone
	c.VehicleRepo = repos.NewVehicleRepository(db)
	c.VehicleTypeRepo = repos.NewVehicleTypeRepository(db)
	c.ZoneRepo = repos.NewZoneRepository(db)
	c.ZoneRateRepo = repos.NewZoneRateRepository(db)
	c.ZoneTypeRepo = repos.NewZoneTypeRepository(db)

	// Customer & Payment
	c.CustomerRepo = repos.NewCustomerRepository(db)
	c.CustomerRegistSourceRepo = repos.NewCustomerRegistSourceRepository(db)
	c.HolidayRepo = repos.NewHolidayRepository(db)
	c.PaymentMethodRepo = repos.NewPaymentMethodRepository(db)

	// Transactional Repos (Directly initialized in struct for brevity if needed)
	c.ParkingTransactionRepo = repos.NewParkingTransactionRepository(db)
	c.TransactionZoneRepo = repos.NewTransactionZoneRepository(db)
	c.TransactionOCRDataRepo = repos.NewTransactionOCRDataRepository(db)
	c.TransactionEventRepo = repos.NewTransactionEventRepository(db)

	// Financials & Logs
	c.PaymentRepo = repos.NewPaymentRepository(db)
	c.RefundRepo = repos.NewRefundRepository(db)
	c.LostTicketFeeRepo = repos.NewLostTicketFeeRepository(db)
}
