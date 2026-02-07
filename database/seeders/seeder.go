package seeders

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"log"
	"parking-management-system-v1/internal/models"
	"time"
)

func RunSeeders(db *gorm.DB) {
	log.Println("Seeding database...")

	// CLEANUP FIRST: Truncate tables in correct order to avoid FK violations
	cleanupDatabase(db)

	// 1. Vehicle Types (Data Master Utama)
	vehicleTypes := []models.VehicleType{
		{Code: "MOTOR", Name: "Motor", Description: "Kendaraan Roda 2", IsActive: true},
		{Code: "CAR", Name: "Mobil", Description: "Kendaraan Roda 4", IsActive: true},
		{Code: "TRUCK", Name: "Truk", Description: "Kendaraan Besar", IsActive: true},
		{Code: "VIP", Name: "VIP", Description: "Kendaraan VIP", IsActive: true},
	}
	for _, v := range vehicleTypes {
		db.FirstOrCreate(&v, models.VehicleType{Code: v.Code})
	}
	log.Println("✓ Vehicle Types seeded")

	// 2. Zone Types
	zoneTypes := []models.ZoneType{
		{Name: "Standard", Description: "Area parkir standar"},
		{Name: "Premium", Description: "Area parkir premium"},
		{Name: "VIP", Description: "Area parkir VIP"},
	}
	for _, zt := range zoneTypes {
		db.FirstOrCreate(&zt, models.ZoneType{Name: zt.Name})
	}
	log.Println("✓ Zone Types seeded")

	// 3. Zones (with occupancy tracking)
	var zoneTypeStandard, zoneTypePremium, zoneTypeVIP models.ZoneType
	db.Where("name = ?", "Standard").First(&zoneTypeStandard)
	db.Where("name = ?", "Premium").First(&zoneTypePremium)
	db.Where("name = ?", "VIP").First(&zoneTypeVIP)

	zones := []models.Zone{
		{
			Name:            "Zone A",
			ZoneTypeID:      zoneTypeStandard.ID,
			Location:        "Lantai 1 Utara",
			MaximumCapacity: 50,
			OccupiedCount:   0,
			AvailableSlots:  50,
			DefaultFee:      2000,
			Is24Hours:       true,
			IsActive:        true,
			IsInMaintenance: false,
		},
		{
			Name:            "Zone B",
			ZoneTypeID:      zoneTypeStandard.ID,
			Location:        "Lantai 1 Selatan",
			MaximumCapacity: 40,
			OccupiedCount:   0,
			AvailableSlots:  40,
			DefaultFee:      2000,
			Is24Hours:       true,
			IsActive:        true,
			IsInMaintenance: false,
		},
		{
			Name:            "VIP Zone",
			ZoneTypeID:      zoneTypeVIP.ID,
			Location:        "Lantai 2 Khusus",
			MaximumCapacity: 20,
			OccupiedCount:   0,
			AvailableSlots:  20,
			DefaultFee:      5000,
			Is24Hours:       true,
			IsActive:        true,
			IsInMaintenance: false,
		},
	}
	for _, z := range zones {
		db.FirstOrCreate(&z, models.Zone{Name: z.Name})
	}
	log.Println("✓ Zones seeded")

	// Fetch zones for later use
	var zoneA, zoneB, zoneVIP models.Zone
	db.Where("name = ?", "Zone A").First(&zoneA)
	db.Where("name = ?", "Zone B").First(&zoneB)
	db.Where("name = ?", "VIP Zone").First(&zoneVIP)

	// 4. Zone Rates (for each vehicle type in each zone)
	var motorType, carType, truckType, vipType models.VehicleType
	db.Where("code = ?", "MOTOR").First(&motorType)
	db.Where("code = ?", "CAR").First(&carType)
	db.Where("code = ?", "TRUCK").First(&truckType)
	db.Where("code = ?", "VIP").First(&vipType)

	zoneRates := []models.ZoneRate{
		// Zone A rates
		{ZoneID: zoneA.ID, VehicleTypeID: motorType.ID, HourlyRate: 2000, DailyMaxRate: 20000, FreeMinutes: 15, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		{ZoneID: zoneA.ID, VehicleTypeID: carType.ID, HourlyRate: 3000, DailyMaxRate: 30000, FreeMinutes: 15, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		{ZoneID: zoneA.ID, VehicleTypeID: truckType.ID, HourlyRate: 5000, DailyMaxRate: 50000, FreeMinutes: 0, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		// Zone B rates
		{ZoneID: zoneB.ID, VehicleTypeID: motorType.ID, HourlyRate: 2000, DailyMaxRate: 20000, FreeMinutes: 15, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		{ZoneID: zoneB.ID, VehicleTypeID: carType.ID, HourlyRate: 3000, DailyMaxRate: 30000, FreeMinutes: 15, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		{ZoneID: zoneB.ID, VehicleTypeID: truckType.ID, HourlyRate: 5000, DailyMaxRate: 50000, FreeMinutes: 0, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		// VIP Zone rates
		{ZoneID: zoneVIP.ID, VehicleTypeID: vipType.ID, HourlyRate: 5000, DailyMaxRate: 50000, FreeMinutes: 30, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
		{ZoneID: zoneVIP.ID, VehicleTypeID: carType.ID, HourlyRate: 5000, DailyMaxRate: 50000, FreeMinutes: 30, IsActive: true, ValidFrom: time.Now().AddDate(0, 0, -1), ValidTo: time.Now().AddDate(1, 0, 0)},
	}
	for _, zr := range zoneRates {
		db.FirstOrCreate(&zr, models.ZoneRate{ZoneID: zr.ZoneID, VehicleTypeID: zr.VehicleTypeID})
	}
	log.Println("✓ Zone Rates seeded")

	// 5. Payment Methods
	paymentMethods := []models.PaymentMethod{
		{Code: "CASH", Name: "Tunai", IsActive: true, Config: pq.StringArray{"no_change:false"}},
		{Code: "QRIS", Name: "QRIS", IsActive: true, Config: pq.StringArray{"provider:gopay", "fee:0"}},
		{Code: "TAPCASH", Name: "Tap Cash", IsActive: true, Config: pq.StringArray{"provider:bca", "fee:0"}},
	}
	for _, p := range paymentMethods {
		db.FirstOrCreate(&p, models.PaymentMethod{Code: p.Code})
	}
	log.Println("✓ Payment Methods seeded")

	// 6. Customers & Vehicles (10 customers, 5 registered with vehicles)
	registeredCount := 0
	for i := 0; i < 10; i++ {
		isRegistered := i < 5 // First 5 are registered
		rfidUID := ""
		if isRegistered {
			rfidUID = gofakeit.UUID()
		}

		customer := models.Customer{
			RfidUID:      rfidUID,
			Name:         gofakeit.Name(),
			Phone:        gofakeit.Phone(),
			IsRegistered: isRegistered,
			TotalVisits:  0,
			TotalSpent:   0,
		}

		if isRegistered {
			now := time.Now()
			customer.RegisteredAt = &now
		}

		db.FirstOrCreate(&customer, models.Customer{RfidUID: customer.RfidUID})

		// If registered, create vehicles
		if isRegistered {
			registeredCount++
			// Fetch the customer again to get the ID
			var createdCustomer models.Customer
			db.Where("rfid_uid = ?", rfidUID).First(&createdCustomer)

			// Create 1-2 vehicles per registered customer
			vehicleCount := 1
			if registeredCount%2 == 0 {
				vehicleCount = 2
			}

			for j := 0; j < vehicleCount; j++ {
				var vehicleTypeID uint
				if j == 0 {
					vehicleTypeID = motorType.ID
				} else {
					vehicleTypeID = carType.ID
				}

				vehicle := models.Vehicle{
					CustomerID:    &createdCustomer.ID,
					VehicleTypeID: vehicleTypeID,
					PlateNumber:   gofakeit.Regex("[A-Z]{1,2}[0-9]{1,4}[A-Z]{1,3}"),
					Description:   gofakeit.Sentence(3),
				}
				db.FirstOrCreate(&vehicle, models.Vehicle{PlateNumber: vehicle.PlateNumber})
			}
		}
	}
	log.Println("✓ Customers & Vehicles seeded (10 customers, 5 registered)")

	// 7. Active Transactions (3 transactions with exit_time IS NULL)
	var registeredCustomers []models.Customer
	db.Where("is_registered = ?", true).Limit(3).Find(&registeredCustomers)

	activeTransactionCount := 0
	for idx, customer := range registeredCustomers {
		if activeTransactionCount >= 3 {
			break
		}

		var vehicle models.Vehicle
		db.Where("customer_id = ?", customer.ID).First(&vehicle)

		var assignedZone models.Zone
		if idx == 0 {
			assignedZone = zoneA
		} else if idx == 1 {
			assignedZone = zoneB
		} else {
			assignedZone = zoneVIP
		}

		entryTime := time.Now().Add(-time.Hour * time.Duration(2+idx))

		transaction := models.ParkingTransaction{
			CustomerID:    &customer.ID,
			OperatorID:    1, // Assuming operator ID 1 exists
			VehicleTypeID: vehicle.VehicleTypeID,
			PlateNumber:   vehicle.PlateNumber,
			RfidUID:       customer.RfidUID,
			EntryTime:     entryTime,
			ExitTime:      nil, // Active transaction
			Status:        "PARKED",
			PaymentStatus: "unpaid",
			Notes:         "Active parking transaction",
		}

		if err := db.Create(&transaction).Error; err == nil {
			// Create transaction zone entry
			transactionZone := models.TransactionZone{
				TransactionID:   transaction.ID,
				SuggestedZoneID: assignedZone.ID,
				ActualZoneID:    assignedZone.ID,
				EntryTime:       entryTime,
				ExitTime:        nil,
			}
			db.Create(&transactionZone)

			// Increment zone occupancy using atomic operation
			db.Model(&models.Zone{}).
				Where("id = ?", assignedZone.ID).
				Updates(map[string]interface{}{
					"occupied_count":  gorm.Expr("occupied_count + ?", 1),
					"available_slots": gorm.Expr("available_slots - ?", 1),
				})

			activeTransactionCount++
		}
	}
	log.Println("✓ Active Transactions seeded (3 transactions)")

	// 8. Historical Transactions (5 completed transactions)
	for i := 0; i < 5; i++ {
		var randomCustomer models.Customer
		db.Where("is_registered = ?", true).Order("RANDOM()").First(&randomCustomer)

		var randomVehicle models.Vehicle
		db.Where("customer_id = ?", randomCustomer.ID).First(&randomVehicle)

		var randomZone models.Zone
		db.Order("RANDOM()").First(&randomZone)

		entryTime := time.Now().AddDate(0, 0, -(5 - i)).Add(-time.Hour * time.Duration(2+i))
		exitTime := entryTime.Add(time.Hour * time.Duration(1+i))
		durationMinutes := int(exitTime.Sub(entryTime).Minutes())

		// Calculate fee (simplified)
		hourlyRate := 3000.0
		durationHours := float64(durationMinutes) / 60.0
		if durationHours < 1 {
			durationHours = 1
		}
		totalFee := hourlyRate * durationHours

		transaction := models.ParkingTransaction{
			CustomerID:      &randomCustomer.ID,
			OperatorID:      1,
			VehicleTypeID:   randomVehicle.VehicleTypeID,
			PlateNumber:     randomVehicle.PlateNumber,
			RfidUID:         randomCustomer.RfidUID,
			EntryTime:       entryTime,
			ExitTime:        &exitTime,
			DurationMinutes: durationMinutes,
			BaseFee:         totalFee,
			TotalFee:        totalFee,
			Status:          "completed",
			PaymentStatus:   "paid",
			Notes:           "Historical completed transaction",
		}

		if err := db.Create(&transaction).Error; err == nil {
			// Create transaction zone entry
			transactionZone := models.TransactionZone{
				TransactionID:   transaction.ID,
				SuggestedZoneID: randomZone.ID,
				ActualZoneID:    randomZone.ID,
				EntryTime:       entryTime,
				ExitTime:        &exitTime,
				DurationMinutes: durationMinutes,
				ZoneFee:         totalFee,
			}
			db.Create(&transactionZone)
		}
	}
	log.Println("✓ Historical Transactions seeded (5 completed)")

	log.Println("✅ Seeding completed successfully!")
}

// cleanupDatabase truncates tables in the correct order to avoid FK violations
func cleanupDatabase(db *gorm.DB) {
	log.Println("Cleaning up existing data...")

	// Order matters: delete child tables first, then parent tables
	db.Exec("TRUNCATE TABLE parking_transactions CASCADE")
	db.Exec("TRUNCATE TABLE transaction_zones CASCADE")
	db.Exec("TRUNCATE TABLE vehicles CASCADE")
	db.Exec("TRUNCATE TABLE customers CASCADE")
	db.Exec("TRUNCATE TABLE zone_rates CASCADE")
	db.Exec("TRUNCATE TABLE zones CASCADE")
	db.Exec("TRUNCATE TABLE zone_types CASCADE")
	db.Exec("TRUNCATE TABLE payment_methods CASCADE")
	db.Exec("TRUNCATE TABLE vehicle_types CASCADE")

	log.Println("✓ Database cleaned")
}
