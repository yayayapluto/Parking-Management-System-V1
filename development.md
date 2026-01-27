# 📋 Module Development Step-by-Step Guide

## 🎯 Development Workflow Overview

```
Phase 1: Foundation & Setup (Week 1)
Phase 2: Core Database & Models (Week 1-2)
Phase 3: Core Business Logic (Week 2-4)
Phase 4: External Integrations (Week 4-5)
Phase 5: Testing & Refinement (Week 5-6)
Phase 6: Deployment & Documentation (Week 6-7)
```

---

# 📦 PHASE 1: Foundation & Setup (Week 1)

## Step 1.1: Initialize Golang Project

**Duration:** 1 hour

### Tasks:

1. Create project root directory `parking-system/`
2. Initialize Go module
3. Create `.gitignore` file
4. Create `.env.example` file
5. Setup `.editorconfig` for code consistency
6. Create basic `README.md`
7. Initialize Git repository
8. Create initial commit

### Files Created:

- `go.mod`
- `go.sum`
- `.gitignore`
- `.env.example`
- `.editorconfig`
- `README.md`

### Dependencies to Install:

- gorilla/mux (HTTP router)
- lib/pq (PostgreSQL driver)
- godotenv (environment variables)
- uber-go/zap (logging)

---

## Step 1.2: Create Project Directory Structure

**Duration:** 30 minutes

### Tasks:

1. Create all main directories:
    - `cmd/`
    - `internal/`
    - `pkg/`
    - `database/migrations/`
    - `config/`
    - `storage/`
    - `tests/`
    - `docs/`
    - `scripts/`
    - `deployments/`
2. Create all subdirectories in `internal/`:
    - `models/`
    - `repositories/`
    - `services/`
    - `handlers/`
    - `dto/requests/`
    - `dto/responses/`
    - `clients/`
    - `circuitbreaker/`
    - `middlewares/`
    - `routes/`
    - `app/`
3. Create all subdirectories in `pkg/`:
    - `config/`
    - `logger/`
    - `errors/`
    - `validator/`
    - `database/`
    - `helpers/`
4. Create placeholder `.gitkeep` files in empty directories

---

## Step 1.3: Setup PostgreSQL Database

**Duration:** 1 hour

### Tasks:

1. Install PostgreSQL locally or setup Docker container
2. Create database `parking_db`
3. Create database user with appropriate permissions
4. Test database connection
5. Install golang-migrate tool for migrations
6. Create migration script templates

### Tools Needed:

- PostgreSQL 15+
- golang-migrate CLI
- psql client (optional)

### Configuration:

- Database name: `parking_db`
- Default user: `postgres`
- Port: `5432`

---

## Step 1.4: Setup Docker Environment

**Duration:** 1-2 hours

### Tasks:

1. Create `Dockerfile` in `deployments/docker/`
2. Create `Dockerfile.dev` for development
3. Create `.dockerignore`
4. Create `docker-compose.yml` with services:
    - PostgreSQL
    - Golang API (parking-system)
    - (Optional) Redis for caching
5. Create `docker-compose.dev.yml` for development
6. Test Docker setup with `docker-compose up`

---

## Step 1.5: Setup Configuration Management

**Duration:** 1 hour

### Tasks:

1. Create `config/config.yaml`
2. Create `config/config.dev.yaml`
3. Create `config/config.prod.yaml`
4. Create `config/config.example.yaml`
5. Create `.env` file from `.env.example`
6. Implement config loading in `pkg/config/config.go`

### Configuration Sections:

- Application settings (name, env, port)
- Database settings
- OCR service settings
- Midtrans settings
- Storage settings
- Camera settings
- JWT settings
- Circuit breaker settings
- Logging settings

---

## Step 1.6: Setup Logging System

**Duration:** 1 hour

### Tasks:

1. Create logger interface in `pkg/logger/logger.go`
2. Choose logging library (zap, logrus, or zerolog)
3. Implement logger initialization
4. Create log levels configuration
5. Setup log file rotation (optional)
6. Test logging in different environments

### Log Levels:

- Debug
- Info
- Warning
- Error
- Fatal

---

## Step 1.7: Setup Makefile

**Duration:** 30 minutes

### Tasks:

1. Create `Makefile` in root directory
2. Add commands:
    - `make help` - Show all commands
    - `make build` - Build application
    - `make run` - Run application
    - `make test` - Run tests
    - `make test-coverage` - Run tests with coverage
    - `make clean` - Clean build artifacts
    - `make migrate-up` - Run migrations up
    - `make migrate-down` - Run migrations down
    - `make docker-build` - Build Docker image
    - `make docker-run` - Run Docker containers
    - `make lint` - Run linter
    - `make fmt` - Format code

---

# 📦 PHASE 2: Core Database & Models (Week 1-2)

## Step 2.1: Create Database Migrations

**Duration:** 2-3 hours

### Migration Files Order:

### Migration 1: vehicles table

- File: `000001_create_vehicles_table.up.sql`
- File: `000001_create_vehicles_table.down.sql`

### Migration 2: zones table

- File: `000002_create_zones_table.up.sql`
- File: `000002_create_zones_table.down.sql`

### Migration 3: parking_transactions table

- File: `000003_create_parking_transactions_table.up.sql`
- File: `000003_create_parking_transactions_table.down.sql`

### Migration 4: payments table

- File: `000004_create_payments_table.up.sql`
- File: `000004_create_payments_table.down.sql`

### Migration 5: ocr_logs table

- File: `000005_create_ocr_logs_table.up.sql`
- File: `000005_create_ocr_logs_table.down.sql`

### Migration 6: zone_rates table

- File: `000006_create_zone_rates_table.up.sql`
- File: `000006_create_zone_rates_table.down.sql`

### Migration 7: operators table

- File: `000007_create_operators_table.up.sql`
- File: `000007_create_operators_table.down.sql`

### Migration 8: zone_occupancy_logs table

- File: `000008_create_zone_occupancy_logs_table.up.sql`
- File: `000008_create_zone_occupancy_logs_table.down.sql`

### Migration 9: system_settings table

- File: `000009_create_system_settings_table.up.sql`
- File: `000009_create_system_settings_table.down.sql`

### Migration 10: indexes and constraints

- File: `000010_add_indexes.up.sql`
- File: `000010_add_indexes_down.sql`

### Tasks:

1. Create each migration file pair (up/down)
2. Define table schemas based on ERD
3. Add primary keys
4. Add foreign keys
5. Add indexes for performance
6. Add unique constraints
7. Add check constraints
8. Add default values
9. Test migrations (up and down)

---

## Step 2.2: Create Model Structs

**Duration:** 2-3 hours

### Files to Create (in order):

### 1. `internal/models/vehicle.go`

- Define Vehicle struct
- Define enums: VehicleType, UserType, RegistrationSource
- Add struct tags for JSON and database
- Add validation tags

### 2. `internal/models/zone.go`

- Define Zone struct
- Define enums: ZoneType, ZoneStatus
- Add struct tags

### 3. `internal/models/parking_transaction.go`

- Define ParkingTransaction struct
- Define enums: ZoneStatus, OCRStatus, PaymentStatus, EntryMode, ExitMode
- Add struct tags
- Add relationships to other models

### 4. `internal/models/payment.go`

- Define Payment struct
- Define enums: PaymentMethod, PaymentStatus
- Add struct tags

### 5. `internal/models/ocr_log.go`

- Define OCRLog struct
- Define enums: OCRType, OCRStatus
- Add struct tags

### 6. `internal/models/zone_rate.go`

- Define ZoneRate struct
- Define enums: DayType
- Add struct tags

### 7. `internal/models/operator.go`

- Define Operator struct
- Define enums: OperatorRole, OperatorStatus
- Add struct tags

### 8. `internal/models/zone_occupancy_log.go`

- Define ZoneOccupancyLog struct
- Define enums: EventType
- Add struct tags

### 9. `internal/models/system_setting.go`

- Define SystemSetting struct
- Define enums: DataType
- Add struct tags

### Common Tasks per Model:

1. Define struct with all fields
2. Add JSON tags
3. Add database tags
4. Add validation tags
5. Define all related enums/constants
6. Add helper methods if needed
7. Document with comments

---

## Step 2.3: Create Database Connection

**Duration:** 1 hour

### Tasks:

1. Create `pkg/database/postgres.go`
2. Implement database connection function
3. Add connection pooling configuration
4. Add connection retry logic
5. Add connection health check
6. Add graceful shutdown
7. Test database connection

### Configuration Parameters:

- Max open connections
- Max idle connections
- Connection max lifetime
- Connection timeout

---

## Step 2.4: Create Repository Interfaces

**Duration:** 2 hours

### Files to Create:

### 1. `internal/repositories/vehicle_repository.go`

Methods needed:
- FindByRFID
- FindByID
- FindByPlate
- Create
- Update
- Delete
- UpdateLastSeen
- IncrementTotalVisits

### 2. `internal/repositories/zone_repository.go`

Methods needed:
- FindByID
- FindAll
- FindAvailable
- FindByType
- UpdateOccupancy
- IncrementOccupied
- DecrementOccupied
- GetAvailability

### 3. `internal/repositories/transaction_repository.go`

Methods needed:
- Create
- FindByID
- FindActiveByRFID
- FindByRFID
- FindPendingZoneEntry
- UpdateZoneStatus
- UpdateOCRStatus
- UpdatePlate
- UpdateExitTime
- UpdatePayment

### 4. `internal/repositories/payment_repository.go`

Methods needed:
- Create
- FindByID
- FindByTransactionID
- FindByExternalID
- UpdateStatus
- UpdatePaymentDetails

### 5. `internal/repositories/ocr_repository.go`

Methods needed:
- Create
- FindByTransactionID
- FindByType

### 6. `internal/repositories/operator_repository.go`

Methods needed:
- Create
- FindByID
- FindByUsername
- Update
- UpdateLastLogin
- IncrementFailedLogins

### 7. `internal/repositories/zone_rate_repository.go`

Methods needed:
- FindByZoneAndType
- FindActiveRates
- Create
- Update

### 8. `internal/repositories/zone_occupancy_repository.go`

Methods needed:
- Create
- FindByZone
- FindByDateRange

### Tasks per Repository:

1. Define interface with all methods
2. Define method signatures with parameters and return types
3. Add comments/documentation
4. Consider pagination needs
5. Consider filtering needs

---

## Step 2.5: Implement Repository Concrete Types

**Duration:** 4-6 hours

### Tasks for Each Repository:

1. Create struct that implements interface
2. Add database connection field
3. Implement constructor (New…Repository)
4. Implement all interface methods with SQL queries
5. Handle errors appropriately
6. Add proper NULL handling
7. Add transaction support where needed
8. Test each method individually

### SQL Query Types Needed:

- SELECT (single row)
- SELECT (multiple rows)
- INSERT
- UPDATE
- DELETE
- Complex JOINs
- Aggregate queries (COUNT, SUM)
- Subqueries

---

# 📦 PHASE 3: Core Business Logic (Week 2-4)

## Step 3.1: Create DTO (Data Transfer Objects)

**Duration:** 2-3 hours

### Request DTOs (in `internal/dto/requests/`):

### Files to Create:

1. `entry_request.go` - Entry gate tap request
2. `zone_select_request.go` - Zone selection request
3. `zone_tap_request.go` - Zone reader tap request
4. `exit_request.go` - Exit gate tap request
5. `exit_confirm_request.go` - Exit confirmation request
6. `payment_request.go` - Payment creation request
7. `vehicle_request.go` - Vehicle CRUD requests
8. `zone_request.go` - Zone CRUD requests
9. `operator_request.go` - Operator CRUD requests
10. `login_request.go` - Authentication request

### Response DTOs (in `internal/dto/responses/`):

### Files to Create:

1. `entry_response.go` - Entry flow responses
2. `zone_availability_response.go` - Zone availability info
3. `exit_response.go` - Exit flow responses
4. `payment_response.go` - Payment info
5. `zone_response.go` - Zone information
6. `vehicle_response.go` - Vehicle information
7. `transaction_response.go` - Transaction details
8. `report_response.go` - Report data
9. `common_response.go` - Standard API response wrapper
10. `error_response.go` - Error response structure

### Tasks per DTO:

1. Define struct with all fields
2. Add JSON tags
3. Add validation tags
4. Add example values in comments
5. Add helper methods (ToModel, FromModel)

---

## Step 3.2: Create Helper Utilities

**Duration:** 2 hours

### Files to Create in `pkg/helpers/`:

### 1. `string.go`

Functions needed:
- GenerateRandomString
- SanitizeString
- TruncateString
- SlugGenerate
- PlateNumberFormat

### 2. `time.go`

Functions needed:
- FormatTimestamp
- ParseTimestamp
- CalculateDuration
- GetCurrentTimestamp
- TimeAgo

### 3. `response.go`

Functions needed:
- SuccessResponse
- ErrorResponse
- PaginationResponse
- ValidationErrorResponse

### 4. `pagination.go`

Functions needed:
- CalculateOffset
- CalculateTotalPages
- BuildPaginationMeta

### 5. `hash.go`

Functions needed:
- HashPassword
- ComparePassword
- GenerateToken

---

## Step 3.3: Create Error Handling System

**Duration:** 1-2 hours

### Files to Create in `pkg/errors/`:

### 1. `errors.go`

Define custom error types:
- AppError (base error)
- ValidationError
- NotFoundError
- ConflictError
- UnauthorizedError
- InternalError

### 2. `codes.go`

Define error codes:
- ErrCodeNotFound
- ErrCodeInvalidInput
- ErrCodeUnauthorized
- ErrCodeConflict
- ErrCodeInternalServer
- etc.

### 3. `handler.go`

Functions needed:
- HandleError
- ConvertToHTTPStatus
- LogError

---

## Step 3.4: Create Validator

**Duration:** 1 hour

### Files to Create in `pkg/validator/`:

### 1. `validator.go`

- Initialize validator instance
- Custom validation rules

### 2. `rules.go`

Custom validators:
- ValidatePlateNumber (Indonesian format)
- ValidateRFIDUID
- ValidatePhoneNumber
- ValidateZoneID

---

## Step 3.5: Create Circuit Breaker

**Duration:** 2-3 hours

### Files to Create in `internal/circuitbreaker/`:

### 1. `breaker.go`

Components:
- State machine (Closed, Open, Half-Open)
- Failure counter
- Timeout handler
- State transitions
- Call wrapper function

### 2. `config.go`

Configuration:
- MaxFailures
- Timeout duration
- Reset timeout
- Per-service configuration

### Tasks:

1. Implement state machine logic
2. Add failure counting
3. Add timeout handling
4. Add state transition logic
5. Add metrics/monitoring hooks
6. Test state transitions

---

## Step 3.6: Create External Service Clients

**Duration:** 3-4 hours

### Files to Create in `internal/clients/`:

### 1. `ocr_client.go`

Methods needed:
- DetectPlate (send image, get plate result)
- HealthCheck
- GetServiceStatus

Features:
- Circuit breaker integration
- Retry logic
- Timeout handling
- Request/response logging

### 2. `midtrans_client.go`

Methods needed:
- GenerateQRIS
- CheckPaymentStatus
- GetTransactionStatus
- HandleCallback

Features:
- Circuit breaker integration
- Signature verification
- Webhook validation
- Error handling

### 3. `storage_client.go`

Methods needed:
- UploadImage (upload to Railway bucket)
- GetImageURL
- DeleteImage
- ListImages

Features:
- File validation
- Size limits
- Format validation
- Error handling

### 4. `camera_client.go`

Methods needed:
- TriggerCapture
- GetSnapshot
- GetCameraStatus

Features:
- Circuit breaker integration
- Timeout handling
- Retry logic

### Tasks per Client:

1. Define client struct
2. Add HTTP client with timeout
3. Integrate circuit breaker
4. Implement all methods
5. Add request/response logging
6. Add error handling
7. Add retry logic
8. Test each method

---

## Step 3.7: Create Service Layer - Vehicle Service

**Duration:** 2 hours

### File: `internal/services/vehicle_service.go`

### Methods to Implement:

1. CreateVehicle
2. GetVehicleByRFID
3. GetVehicleByID
4. GetVehicleByPlate
5. UpdateVehicle
6. DeleteVehicle
7. SearchVehicles
8. GetVehicleStatistics

### Business Logic:

- Validate RFID uniqueness
- Validate plate number format
- Auto-assign vehicle type
- Update last seen timestamp
- Track total visits

---

## Step 3.8: Create Service Layer - Zone Service

**Duration:** 2 hours

### File: `internal/services/zone_service.go`

### Methods to Implement:

1. GetAllZones
2. GetZoneByID
3. GetAvailableZones
4. GetZoneAvailability
5. UpdateZoneCapacity
6. MarkZoneFull
7. MarkZoneAvailable
8. GetZoneOccupancyLog

### Business Logic:

- Calculate available slots
- Determine zone status (active, full, maintenance)
- Suggest zones based on vehicle type
- Real-time capacity tracking

---

## Step 3.9: Create Service Layer - Entry Service

**Duration:** 3-4 hours

### File: `internal/services/entry_service.go`

### Methods to Implement:

1. ProcessEntryTap (RFID tap at entry gate)
2. CheckUserType (registered/guest)
3. CheckActiveParking
4. GetAvailableZones
5. SuggestZone
6. SelectZone
7. CreateTransaction
8. TriggerOCRProcess (async)

### Business Logic Flow:

1. Read RFID UID
2. Check in database (registered or guest)
3. If guest → auto-create record
4. Check if already has active parking
5. If active → reject
6. Query available zones
7. Suggest zone based on vehicle type
8. User selects zone
9. Create parking transaction
10. Trigger camera capture (background)
11. Return ticket info

---

## Step 3.10: Create Service Layer - Zone Entry Service

**Duration:** 2 hours

### File: `internal/services/zone_entry_service.go`

### Methods to Implement:

1. ProcessZoneTap (RFID tap at zone reader)
2. ValidatePendingTransaction
3. ValidateZoneAssignment
4. CheckZoneCapacity
5. OccupyZone (atomic transaction)
6. LogZoneEntry

### Business Logic Flow:

1. Read RFID UID at zone reader
2. Find pending transaction
3. Validate zone matches assignment
4. Check zone capacity
5. Begin database transaction
6. Update zone occupancy (+1)
7. Update transaction (zone_status = occupied)
8. Log zone occupancy
9. Commit transaction
10. Return success

---

## Step 3.11: Create Service Layer - Fee Calculator

**Duration:** 2 hours

### File: `internal/services/fee_calculator.go`

### Methods to Implement:

1. CalculateFee (based on duration and zone)
2. GetZoneRate (get applicable rate)
3. CalculateDuration
4. ApplyTimeBasedRules (peak hours, happy hours)
5. ApplyDiscounts (membership, promo)
6. RoundFee

### Business Logic:

- Get zone entry time and current time
- Calculate duration in minutes/hours
- Get applicable zone rate (vehicle type, day type, time)
- Apply time-based multipliers (surge pricing)
- Apply discounts if applicable
- Round to nearest hundred/thousand
- Return fee breakdown

---

## Step 3.12: Create Service Layer - Exit Service

**Duration:** 3-4 hours

### File: `internal/services/exit_service.go`

### Methods to Implement:

1. ProcessExitTap (RFID tap at exit gate)
2. FindActiveTransaction
3. GetTransactionDetails
4. ConfirmExit
5. CalculateFeeForTransaction
6. GenerateReceipt
7. ReleaseZone
8. UpdateVehicleStatistics
9. TriggerExitOCR (async)

### Business Logic Flow:

1. Read RFID UID at exit gate
2. Find active transaction (zone_status = occupied)
3. Get transaction details (entry time, zone, plate, photo)
4. Display info to user
5. User confirms exit
6. Calculate parking fee
7. Generate payment (QRIS or manual)
8. Wait for payment confirmation
9. Release zone occupancy (-1)
10. Update transaction (exit_time, fee, status)
11. Update vehicle statistics (last_seen, total_visits)
12. Generate receipt
13. Trigger exit photo OCR (background)

---

## Step 3.13: Create Service Layer - Payment Service

**Duration:** 3-4 hours

### File: `internal/services/payment_service.go`

### Methods to Implement:

1. CreatePayment
2. GenerateQRIS (via Midtrans)
3. ProcessManualPayment
4. HandlePaymentCallback (webhook)
5. CheckPaymentStatus
6. UpdatePaymentStatus
7. ProcessRefund
8. GenerateReceiptNumber

### Business Logic:

- Create payment record
- Generate QRIS via Midtrans client
- Store QR code URL and external ID
- Set expiry time (5 minutes)
- Handle webhook callback from Midtrans
- Verify signature
- Update payment status
- Trigger zone release if paid

---

## Step 3.14: Create Service Layer - OCR Service

**Duration:** 2-3 hours

### File: `internal/services/ocr_service.go`

### Methods to Implement:

1. ProcessEntryOCR (background job)
2. ProcessExitOCR (background job)
3. CallOCRService (via OCR client)
4. ValidatePlateFormat
5. HandleOCRResult
6. LogOCRAttempt
7. CheckPlateMismatch (registered user)
8. AutoRegisterGuestVehicle (optional)
9. CompareEntryExitPlates (security check)

### Business Logic Flow (Entry OCR):

1. Receive transaction ID and image data
2. Call Python OCR service (with circuit breaker)
3. Receive OCR result (plate, confidence)
4. Validate plate format
5. Check user type
    - If registered → compare with DB plate
    - If mismatch → log warning
    - If guest → update transaction with OCR plate
6. Optionally auto-register guest vehicle
7. Log OCR result
8. Update transaction OCR status

### Business Logic Flow (Exit OCR):

1. Capture exit photo
2. Call Python OCR service
3. Compare exit plate with entry plate
4. If mismatch → security alert
5. Log OCR result

---

## Step 3.15: Create Service Layer - Receipt Generator

**Duration:** 1-2 hours

### File: `internal/services/receipt_generator.go`

### Methods to Implement:

1. GenerateReceipt (text format)
2. GenerateQRCodeForReceipt
3. FormatReceiptData
4. GetReceiptTemplate

### Receipt Content:

- Transaction ID
- RFID UID
- Plate number
- Zone
- Entry time
- Exit time
- Duration
- Fee
- Payment method
- Payment status
- QR code (for digital receipt)

---

## Step 3.16: Create Service Layer - Report Service

**Duration:** 2-3 hours

### File: `internal/services/report_service.go`

### Methods to Implement:

1. GetDailyReport
2. GetMonthlyReport
3. GetZoneReport
4. GetRevenueReport
5. GetOccupancyTrends
6. GetPeakHours
7. GetTopVehicles

### Report Types:

- Daily transaction summary
- Monthly revenue
- Zone utilization
- Peak hour analysis
- Vehicle frequency
- Payment method breakdown

---

## Step 3.17: Create Service Layer - Security Service

**Duration:** 1-2 hours

### File: `internal/services/security_service.go`

### Methods to Implement:

1. GetPlateMismatches
2. GetSecurityAlerts
3. FlagSuspiciousTransaction
4. GetAbandonedVehicles (parked > 24h)
5. BlacklistVehicle
6. CheckBlacklist

---

# 📦 PHASE 4: HTTP Layer & Routing (Week 4)

## Step 4.1: Create Middlewares

**Duration:** 2-3 hours

### Files to Create in `internal/middlewares/`:

### 1. `logger.go`

- Log HTTP requests
- Log response status
- Log duration
- Log request ID

### 2. `cors.go`

- Configure CORS headers
- Allow origins
- Allow methods
- Allow headers

### 3. `error_handler.go`

- Catch panics
- Convert errors to HTTP responses
- Log errors

### 4. `auth.go`

- Verify JWT token
- Extract operator info
- Check permissions

### 5. `rate_limiter.go`

- Limit requests per IP
- Limit requests per user
- Prevent abuse

### 6. `request_id.go`

- Generate unique request ID
- Add to context
- Add to response headers

---

## Step 4.2: Create Handlers - Vehicle Handler

**Duration:** 2 hours

### File: `internal/handlers/vehicle_handler.go`

### Endpoints to Implement:

1. Create - POST /api/v1/vehicles
2. GetByRFID - GET /api/v1/vehicles/:rfid_uid
3. Update - PUT /api/v1/vehicles/:id
4. Delete - DELETE /api/v1/vehicles/:id
5. Search - GET /api/v1/vehicles/search

### Handler Responsibilities:

- Parse request body/params
- Validate input
- Call service layer
- Format response
- Handle errors

---

## Step 4.3: Create Handlers - Zone Handler

**Duration:** 1-2 hours

### File: `internal/handlers/zone_handler.go`

### Endpoints to Implement:

1. GetAll - GET /api/v1/zones
2. GetByID - GET /api/v1/zones/:zone_id
3. GetAvailability - GET /api/v1/zones/:zone_id/availability
4. Create - POST /api/v1/zones
5. Update - PUT /api/v1/zones/:zone_id

---

## Step 4.4: Create Handlers - Entry Handler

**Duration:** 2-3 hours

### File: `internal/handlers/entry_handler.go`

### Endpoints to Implement:

1. TapEntry - POST /api/v1/parking/entry/tap
2. SelectZone - POST /api/v1/parking/entry/select-zone
3. TapZone - POST /api/v1/parking/zone/tap

### Request Flow:

Entry Tap → Select Zone → Zone Tap

---

## Step 4.5: Create Handlers - Exit Handler

**Duration:** 2-3 hours

### File: `internal/handlers/exit_handler.go`

### Endpoints to Implement:

1. TapExit - POST /api/v1/parking/exit/tap
2. ConfirmExit - POST /api/v1/parking/exit/:transaction_id/confirm
3. GenerateQRIS - POST /api/v1/parking/exit/:transaction_id/generate-qris
4. GetPaymentStatus - GET /api/v1/parking/exit/:transaction_id/payment-status

---

## Step 4.6: Create Handlers - Payment Handler

**Duration:** 1-2 hours

### File: `internal/handlers/payment_handler.go`

### Endpoints to Implement:

1. CreatePayment - POST /api/v1/payments
2. GetPaymentStatus - GET /api/v1/payments/:payment_id/status
3. GetPaymentHistory - GET /api/v1/payments/history

---

## Step 4.7: Create Handlers - Webhook Handler

**Duration:** 2 hours

### File: `internal/handlers/webhook_handler.go`

### Endpoints to Implement:

1. MidtransCallback - POST /webhooks/midtrans/notification

### Webhook Responsibilities:

- Verify signature
- Parse notification
- Validate request
- Call payment service
- Return success response

---

## Step 4.8: Create Handlers - Report Handler

**Duration:** 1-2 hours

### File: `internal/handlers/report_handler.go`

### Endpoints to Implement:

1. GetDailyReport - GET /api/v1/reports/daily
2. GetMonthlyReport - GET /api/v1/reports/monthly
3. GetZoneReport - GET /api/v1/reports/zone/:zone_id
4. GetRevenueReport - GET /api/v1/reports/revenue

---

## Step 4.9: Create Handlers - Operator Handler

**Duration:** 1-2 hours

### File: `internal/handlers/operator_handler.go`

### Endpoints to Implement:

1. Login - POST /api/v1/auth/login
2. Logout - POST /api/v1/auth/logout
3. GetProfile - GET /api/v1/operators/me
4. CreateOperator - POST /api/v1/operators
5. UpdateOperator - PUT /api/v1/operators/:id

---

## Step 4.10: Create Routes

**Duration:** 2 hours

### File: `internal/routes/routes.go`

### Route Groups:

1. Public routes (no auth):
    - Health check
    - Webhook
2. Parking routes:
    - Entry flow
    - Zone entry
    - Exit flow
3. Management routes (auth required):
    - Vehicles CRUD
    - Zones CRUD
    - Reports
4. Admin routes (admin only):
    - Operators CRUD
    - System settings

### Tasks:

1. Setup router (gorilla/mux)
2. Apply global middlewares
3. Group routes by functionality
4. Apply specific middlewares per route
5. Register all handlers
6. Add route documentation

---

# 📦 PHASE 5: Application Bootstrap (Week 4-5)

## Step 5.1: Create Container (Dependency Injection)

**Duration:** 2-3 hours

### 

File: `internal/app/container.go`

### Components to Initialize:

1. Configuration
2. Logger
3. Database connection
4. Repositories (all)
5. Clients (OCR, Midtrans, Storage, Camera)
6. Services (all)
7. Handlers (all)

### Container Structure:

- Config
- Logger
- DB
- All repositories
- All clients
- All services
- All handlers

### Methods:

- NewContainer (initialize all dependencies)
- initDatabase
- initClients
- initRepositories
- initServices
- initHandlers
- Close (cleanup)

---

## Step 5.2: Create Application

**Duration:** 1-2 hours

### File: `internal/app/app.go`

### Application Struct:

- Container
- HTTP Server

### Methods:

- New (create app instance)
- Run (start application)
- Shutdown (graceful shutdown)

### Run Flow:

1. Load configuration
2. Initialize container
3. Setup routes
4. Create HTTP server
5. Start server in goroutine
6. Wait for interrupt signal
7. Graceful shutdown
8. Close all connections

---

## Step 5.3: Create Database Helper

**Duration:** 1 hour

### File: `internal/app/database.go`

### Functions:

- InitDatabase (connect to PostgreSQL)
- PingDatabase (health check)
- CloseDatabase (close connection)
- RunMigrations (optional)

---

## Step 5.4: Create Server Helper

**Duration:** 1 hour

### File: `internal/app/server.go`

### Functions:

- NewHTTPServer (create HTTP server)
- ConfigureServer (timeouts, limits)
- StartServer
- ShutdownServer (graceful)

---

## Step 5.5: Update Main Entry Point

**Duration:** 30 minutes

### File: `cmd/main.go`

### Tasks:

1. Import app package
2. Call app.New()
3. Call app.Run()
4. Handle errors

---

# 📦 PHASE 6: Python OCR Service (Week 5)

## Step 6.1: Initialize Python Project

**Duration:** 1 hour

### Tasks:

1. Create `ocr-service/` directory
2. Create virtual environment
3. Create `requirements.txt`
4. Create `requirements-dev.txt`
5. Create `.gitignore`
6. Create `README.md`
7. Create directory structure

---

## Step 6.2: Setup FastAPI Application

**Duration:** 1 hour

### Files to Create:

1. `app/main.py` - Main FastAPI app
2. `app/__init__.py`
3. `run.py` - Run server script

---

## Step 6.3: Create OCR Routers

**Duration:** 1 hour

### Files to Create:

1. `app/routers/ocr.py` - OCR endpoints
2. `app/routers/health.py` - Health check

### Endpoints:

- POST /ocr/plate
- GET /health

---

## Step 6.4: Create Image Preprocessor

**Duration:** 2-3 hours

### File: `app/services/preprocessor.py`

### Functions Needed:

1. ValidateImage (format, size)
2. DecodeBase64Image
3. ResizeImage
4. GrayscaleConversion
5. NoiseReduction
6. ContrastEnhancement
7. EdgeDetection
8. DeskewImage

---

## Step 6.5: Implement OCR Engines

**Duration:** 4-6 hours

### Files to Create:

### 1. `app/services/easy_ocr_engine.py`

- Initialize EasyOCR reader
- Detect text in image
- Return results

### 2. `app/services/paddle_ocr_engine.py`

- Initialize PaddleOCR
- Detect text in image
- Return results

### 3. `app/services/tesseract_engine.py` (optional)

- Initialize Tesseract
- Detect text in image
- Return results

---

## Step 6.6: Create Ensemble Service

**Duration:** 2-3 hours

### File: `app/services/ensemble.py`

### Functions:

1. RunMultipleOCR (run all engines)
2. AggregateResults (voting mechanism)
3. SelectBestResult (highest confidence)
4. MergeResults

---

## Step 6.7: Create Plate Validator

**Duration:** 2 hours

### File: `app/validators/plate_validator.py`

### Functions:

1. ValidateIndonesianPlateFormat
2. CleanPlateText (remove spaces, special chars)
3. NormalizePlateText
4. CheckPlatePattern

### Indonesian Plate Patterns:

- B 1234 XYZ
- D 1234 XYZ
- AB 1234 XYZ

---

## Step 6.8: Create OCR Service

**Duration:** 2 hours

### File: `app/services/ocr_service.py`

### Main OCR Flow:

1. Receive base64 image
2. Validate image
3. Preprocess image
4. Run OCR engines (ensemble)
5. Aggregate results
6. Validate plate format
7. Return best result with confidence

---

## Step 6.9: Create Request/Response Models

**Duration:** 1 hour

### Files:

1. `app/models/request.py`
2. `app/models/response.py`

---

## Step 6.10: Create Docker for OCR Service

**Duration:** 1 hour

### Files to Create:

1. `Dockerfile`
2. `docker-compose.yml`
3. `.dockerignore`

---

# 📦 PHASE 7: Testing (Week 5-6)

## Step 7.1: Unit Tests - Repositories

**Duration:** 3-4 hours

### Files to Create in `tests/unit/repositories/`:

1. `vehicle_repository_test.go`
2. `zone_repository_test.go`
3. `transaction_repository_test.go`

### Test Cases per Repository:

- Test Create
- Test FindByID
- Test Update
- Test Delete
- Test edge cases
- Test error handling

---

## Step 7.2: Unit Tests - Services

**Duration:** 4-6 hours

### Files to Create in `tests/unit/services/`:

1. `entry_service_test.go`
2. `exit_service_test.go`
3. `payment_service_test.go`
4. `fee_calculator_test.go`

### Testing Strategy:

- Use mock repositories
- Test business logic
- Test error scenarios
- Test edge cases

---

## Step 7.3: Create Mock Objects

**Duration:** 2-3 hours

### Files to Create in `tests/mocks/`:

1. `mock_vehicle_repository.go`
2. `mock_zone_repository.go`
3. `mock_transaction_repository.go`
4. `mock_ocr_client.go`
5. `mock_midtrans_client.go`

---

## Step 7.4: Integration Tests

**Duration:** 3-4 hours

### Files to Create in `tests/integration/`:

1. `entry_flow_test.go` - Test complete entry flow
2. `exit_flow_test.go` - Test complete exit flow
3. `payment_flow_test.go` - Test payment process

### Setup:

- Use test database
- Seed test data
- Clean up after tests

---

## Step 7.5: E2E Tests

**Duration:** 2-3 hours

### Files to Create in `tests/e2e/`:

1. `complete_parking_flow_test.go`
2. `concurrent_parking_test.go`

### Scenarios:

- Complete parking cycle (entry → park → exit → pay)
- Multiple concurrent users
- Edge cases (zone full, payment timeout)

---

## Step 7.6: Load Testing

**Duration:** 2 hours

### Tools:

- k6 or Apache JMeter

### Scenarios:

- 100 concurrent users entering
- 50 concurrent payments
- Database connection pool stress test

---

# 📦 PHASE 8: Deployment & Documentation (Week 6-7)

## Step 8.1: Prepare Deployment Files

**Duration:** 2-3 hours

### Files to Create:

### 1. Railway Deployment

- `railway.toml`
- `railway.json`

### 2. Kubernetes (optional)

- `deployments/kubernetes/deployment.yaml`
- `deployments/kubernetes/service.yaml`
- `deployments/kubernetes/configmap.yaml`
- `deployments/kubernetes/secret.yaml`

### 3. Docker Production

- `Dockerfile` (optimized)
- `docker-compose.prod.yml`

---

## Step 8.2: Create API Documentation

**Duration:** 3-4 hours

### Files to Create:

### 1. OpenAPI Specification

- `docs/api/openapi.yaml`

### 2. Postman Collection

- `docs/api/postman_collection.json`

### 3. API Documentation

- `docs/api/api_documentation.md`

### Content:

- All endpoints
- Request/response examples
- Authentication
- Error codes
- Rate limits

---

## Step 8.3: Create Architecture Documentation

**Duration:** 2-3 hours

### Files to Create in `docs/architecture/`:

### 1. `system_design.md`

- Architecture overview
- Component diagram
- Data flow diagram

### 2. `database_schema.md`

- ERD
- Table descriptions
- Relationships

### 3. `flow_diagrams.md`

- Entry flow
- Exit flow
- Payment flow

### 4. `deployment.md`

- Deployment architecture
- Infrastructure requirements
- Scaling strategy

---

## Step 8.4: Create Guides

**Duration:** 3-4 hours

### Files to Create in `docs/guides/`:

### 1. `setup_guide.md`

- Prerequisites
- Installation steps
- Configuration
- Running locally

### 2. `developer_guide.md`

- Project structure
- Development workflow
- Code standards
- Testing guidelines

### 3. `deployment_guide.md`

- Environment setup
- Database migrations
- Secrets management
- Deploy to Railway/K8s

### 4. `troubleshooting.md`

- Common issues
- Solutions
- Debug tips

---

## Step 8.5: Create README

**Duration:** 1-2 hours

### Sections:

1. Project Overview
2. Features
3. Tech Stack
4. Architecture
5. Prerequisites
6. Installation
7. Configuration
8. Usage
9. API Documentation
10. Development
11. Testing
12. Deployment
13. Contributing
14. License

---

## Step 8.6: Environment Configuration

**Duration:** 1 hour

### Tasks:

1. Create `.env.example` with all variables
2. Create `.env.dev` for development
3. Create `.env.prod` for production
4. Document all environment variables
5. Setup secrets in Railway/K8s

---

## Step 8.7: Database Seeding

**Duration:** 1-2 hours

### Create Seed Scripts:

1. `scripts/seed.sh`
2. Seed data:
    - Default zones (A, B, C, VIP)
    - Zone rates
    - Default operator account
    - System settings

---

## Step 8.8: Monitoring & Logging

**Duration:** 2-3 hours

### Setup:

1. Configure structured logging
2. Setup log aggregation (optional)
3. Setup metrics collection (Prometheus)
4. Setup alerting (critical errors)
5. Setup health check endpoints

---

## Step 8.9: Security Hardening

**Duration:** 2-3 hours

### Tasks:

1. Enable HTTPS
2. Configure CORS properly
3. Add rate limiting
4. Input sanitization
5. SQL injection prevention
6. Secrets management
7. JWT token security
8. API key rotation

---

## Step 8.10: Performance Optimization

**Duration:** 2-3 hours

### Tasks:

1. Add database indexes
2. Optimize queries
3. Add caching (Redis - optional)
4. Connection pooling tuning
5. Image optimization
6. Response compression

---

## Step 8.11: Final Testing & QA

**Duration:** 2-3 hours

### Checklist:

- ✅ All unit tests passing
- ✅ All integration tests passing
- ✅ E2E tests passing
- ✅ Load tests acceptable
- ✅ Security scan clean
- ✅ API documentation complete
- ✅ Deployment successful
- ✅ Health checks working
- ✅ Logs properly formatted
- ✅ Error handling working

---

## Step 8.12: Deploy to Production

**Duration:** 2-4 hours

### Deployment Steps:

1. Backup current database (if migration)
2. Run database migrations
3. Deploy application
4. Run smoke tests
5. Monitor logs
6. Verify all endpoints
7. Test critical flows
8. Enable monitoring/alerting

---

# 📊 Total Estimated Time

| Phase | Duration |
| --- | --- |
| Phase 1: Foundation & Setup | 1 week |
| Phase 2: Database & Models | 1 week |
| Phase 3: Business Logic | 2 weeks |
| Phase 4: HTTP Layer | 1 week |
| Phase 5: Bootstrap | 2-3 days |
| Phase 6: Python OCR | 1 week |
| Phase 7: Testing | 1 week |
| Phase 8: Deployment | 1 week |
| **TOTAL** | **6-7 weeks** |

---

# 🎯 Daily Development Checklist Template

## Daily Tasks:

- [ ]  Review yesterday’s work
- [ ]  Write tests first (TDD)
- [ ]  Implement feature
- [ ]  Run tests
- [ ]  Code review (self or peer)
- [ ]  Update documentation
- [ ]  Commit with meaningful message
- [ ]  Push to repository

---

# ✅ Module Completion Checklist

## Per Module:

- [ ]  Interface/struct defined
- [ ]  All methods implemented
- [ ]  Unit tests written
- [ ]  Unit tests passing
- [ ]  Integration test written (if applicable)
- [ ]  Documentation added
- [ ]  Error handling implemented
- [ ]  Logging added
- [ ]  Code reviewed
- [ ]  Committed to Git

---