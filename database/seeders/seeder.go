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

	// 1. Vehicle Types (Data Master Utama)
	vehicleTypes := []models.VehicleType{
		{Code: "CAR", Name: "Mobil", Description: "Kendaraan Roda 4", IsActive: true},
		{Code: "MOTOR", Name: "Motor", Description: "Kendaraan Roda 2", IsActive: true},
		{Code: "TRUCK", Name: "Truk/Bus", Description: "Kendaraan Besar", IsActive: true},
	}
	for _, v := range vehicleTypes {
		db.FirstOrCreate(&v, models.VehicleType{Code: v.Code})
	}

	// 2. Customer Registration Sources (Random Data)
	for i := 0; i < 5; i++ {
		source := models.CustomerRegistSource{
			Name:        gofakeit.Company(),
			Description: gofakeit.Sentence(5),
		}
		db.FirstOrCreate(&source, models.CustomerRegistSource{Name: source.Name})
	}

	// 3. Payment Methods (Keperluan Fintech)
	paymentMethods := []models.PaymentMethod{
		{Code: "CASH", Name: "Tunai", IsActive: true, Config: pq.StringArray{"no_change:false"}},
		{Code: "QRIS", Name: "QRIS", IsActive: true, Config: pq.StringArray{"provider:gopay", "fee:0"}},
		{Code: "BANK_TRANSFER", Name: "Transfer Bank", IsActive: true, Config: pq.StringArray{"bank:bca", "admin_fee:0"}},
	}
	for _, p := range paymentMethods {
		db.FirstOrCreate(&p, models.PaymentMethod{Code: p.Code})
	}

	// 4. Holidays (Random Dates)
	for i := 0; i < 10; i++ {
		// Generate tanggal acak dalam rentang 1 tahun ke depan
		randomDate := gofakeit.DateRange(time.Now(), time.Now().AddDate(1, 0, 0))

		holiday := models.Holiday{
			Date:         randomDate,
			Name:         gofakeit.Adjective() + " Holiday",
			AffectsRates: gofakeit.Float64Range(1.1, 2.0),
		}
		// Gunakan FirstOrCreate berdasarkan tanggal agar tidak duplikat
		db.FirstOrCreate(&holiday, models.Holiday{Date: holiday.Date})
	}

	// 5. Zone Types (Data Wilayah Parkir)
	// Kita buat beberapa zona default dulu
	defaultZones := []models.ZoneType{
		{Name: "Main Area", Description: "Area parkir utama dekat pintu masuk"},
		{Name: "Basement 1", Description: "Area parkir bawah tanah lantai 1"},
		{Name: "VIP Zone", Description: "Area parkir khusus member VIP"},
	}
	for _, z := range defaultZones {
		db.FirstOrCreate(&z, models.ZoneType{Name: z.Name})
	}

	// Tambah random zone pake Faker
	for i := 0; i < 3; i++ {
		zone := models.ZoneType{
			Name:        gofakeit.Adjective() + " Sector",
			Description: gofakeit.Sentence(5),
		}
		db.FirstOrCreate(&zone, models.ZoneType{Name: zone.Name})
	}

	log.Println("Seeding completed successfully.")
}
