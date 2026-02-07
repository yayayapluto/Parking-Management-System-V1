# Parking Exit & Fee Calculation Implementation

## Overview

This document describes the implementation of the Parking Exit (Gate Out) service and handler, including a dynamic fee
calculation engine for the Parking Management System.

## Architecture

### Components

1. **Fee Calculator Service** (`internal/services/fee_calculator.go`)
    - Calculates parking fees based on duration and zone rates
    - Implements hourly rate calculation with rounding up logic
    - Applies daily maximum rates when applicable

2. **Parking Transaction Service** (`internal/services/parking_transaction.go`)
    - Handles vehicle exit logic with atomic database transactions
    - Manages zone occupancy updates
    - Updates vehicle and customer statistics

3. **Parking Transaction Handler** (`internal/handlers/parking_transaction_handler.go`)
    - HTTP endpoint handler for exit requests
    - Validates incoming requests
    - Returns formatted exit response with receipt details

4. **DTOs** (`internal/dto/requests/parking_transaction.go`, `internal/dto/responses/parking_transaction.go`)
    - `ParkingExitRequest`: Contains RFID UID and payment method
    - `ParkingExitResponse`: Contains transaction details for receipt printing

## API Specification

### Endpoint

```
POST /api/v1/transactions/exit
```

### Request

```json
{
  "rfid_uid": "string (required, max 50 chars)",
  "payment_method": "string (required, enum: manual|qris)"
}
```

### Response (Success - 200 OK)

```json
{
  "success": true,
  "message": "Vehicle exit recorded successfully",
  "data": {
    "transaction_id": "hashed_id",
    "rfid_uid": "RFID123456",
    "plate_number": "B 1234 ABC",
    "entry_time": "2026-02-07T10:00:00Z",
    "exit_time": "2026-02-07T12:30:00Z",
    "total_duration": 150,
    "total_fee": 30000.00,
    "payment_status": "paid",
    "payment_method": "manual"
  }
}
```

### Response (Error - 404 Not Found)

```json
{
  "success": false,
  "message": "no active parking found",
  "error": "no active parking found for the given RFID"
}
```

## Database Logic

### Transaction Lookup

The system finds the active parking transaction using:

- `rfid_uid` = provided RFID UID
- `exit_time IS NULL` (not yet exited)
- `status = 'PARKED'` (currently parked)

### Duration Calculation

Duration is calculated as:

```
duration_minutes = (exit_time - entry_time) in minutes
duration_hours = CEIL(duration_minutes / 60)
minimum_charge = 1 hour
```

Example:

- 30 minutes → 1 hour charge
- 2 hours 30 minutes → 3 hours charge
- 2 hours → 2 hours charge

### Fee Calculation Engine

The fee calculation follows this logic:

1. **Fetch Zone Rate**: Query `zone_rates` table for:
    - Matching `zone_id` and `vehicle_type_id`
    - `is_active = true`
    - Valid date range: `valid_from <= today AND valid_to >= today`
    - Ordered by `effective_hour_from` for time-specific rates

2. **Calculate Base Fee**:
   ```
   base_fee = hourly_rate * duration_hours
   ```

3. **Apply Daily Maximum**:
   ```
   if daily_max_rate > 0 AND base_fee > daily_max_rate:
     total_fee = daily_max_rate
   else:
     total_fee = base_fee
   ```

### Atomic Database Transaction

All updates are performed within a single database transaction to ensure data consistency:

1. **Update `parking_transactions`**:
    - Set `exit_time = NOW`
    - Set `duration_minutes = calculated_duration`
    - Set `base_fee = calculated_base_fee`
    - Set `total_fee = calculated_total_fee`
    - Set `status = 'completed'`
    - Set `payment_status = 'paid'`

2. **Update `transaction_zones`**:
    - Set `exit_time = NOW`
    - Set `duration_minutes = calculated_duration`
    - Set `zone_fee = calculated_total_fee`

3. **Update `zones`**:
    - Decrease `occupied_count` by 1
    - Increase `available_slots` by 1
    - Update `last_exit_time = NOW`

4. **Update Vehicle Stats** (if customer is registered):
    - Update `vehicles.last_seen_at = NOW`
    - Increment `customers.total_visits` by 1

## Safety & Edge Cases

### No Active Parking

If no active parking transaction is found:

- Returns HTTP 404 Not Found
- Error message: "no active parking found for the given RFID"

### Payment Failure Handling

The current implementation marks transactions as paid immediately. For payment failure scenarios:

- The transaction remains in the database with `payment_status = 'paid'`
- Future enhancements should implement payment gateway integration
- Implement retry logic and timeout handling

### Concurrency Control

The implementation uses GORM's atomic transactions to prevent race conditions:

- Database-level locking ensures only one exit per transaction
- Zone occupancy updates are atomic
- Customer statistics updates are atomic using `gorm.Expr("total_visits + ?", 1)`

### Validation

Input validation includes:

- RFID UID: Required, max 50 characters
- Payment Method: Required, must be "manual" or "qris"

## Implementation Details

### Fee Calculator (`FeeCalculator` interface)

```go
type FeeCalculator interface {
CalculateFeeWithDB(
ctx context.Context,
entryTime time.Time,
exitTime time.Time,
zoneID uint,
vehicleTypeID uint,
db *gorm.DB,
) (FeeCalculationResult, error)
}
```

### Service Dependencies

The `ParkingTransactionService` requires:

- `ParkingTransactionRepository`: For transaction CRUD operations
- `ZoneRepository`: For zone information
- `ZoneRateRepository`: For rate lookups
- `VehicleRepository`: For vehicle updates
- `CustomerRepository`: For customer updates
- `IDObfuscator`: For ID encoding/decoding
- `Validator`: For request validation
- `FeeCalculator`: For fee calculations

### Container Initialization

The service is initialized in `internal/container/services_container.go`:

```go
feeCalculator := services.NewFeeCalculator(c.ZoneRateRepo)

c.ParkingTransactionService = services.NewParkingTransactionService(
c.ParkingTransactionRepo,
c.ZoneRepo,
c.ZoneRateRepo,
c.VehicleRepo,
c.CustomerRepo,
o,
v,
feeCalculator,
)
```

## Testing Scenarios

### Scenario 1: Successful Exit with Manual Payment

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID123456",
    "payment_method": "manual"
  }'
```

Expected: 200 OK with transaction details

### Scenario 2: No Active Parking

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "NONEXISTENT",
    "payment_method": "manual"
  }'
```

Expected: 404 Not Found with error message

### Scenario 3: Invalid Payment Method

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID123456",
    "payment_method": "invalid"
  }'
```

Expected: 400 Bad Request with validation error

## Future Enhancements

1. **Payment Gateway Integration**: Implement actual payment processing with QRIS and manual payment methods
2. **Discount Handling**: Add support for membership discounts and promotional codes
3. **Lost Ticket Fee**: Implement lost ticket fee calculation for vehicles without entry records
4. **Multi-Zone Parking**: Handle vehicles that move between zones during parking
5. **Refund Processing**: Implement refund logic for overpayments
6. **Audit Logging**: Add comprehensive audit trails for all transactions
7. **Real-time Notifications**: Send SMS/email notifications on exit
8. **Receipt Generation**: Integrate with receipt printer for physical tickets

## Notes

- OCR integration is not included in this implementation (handled by separate Python service)
- The implementation assumes a single zone per parking session
- Payment processing is simplified and should be enhanced for production use
- All timestamps are in UTC and should be converted to local timezone for display
