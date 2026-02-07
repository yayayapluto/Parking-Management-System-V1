# Comprehensive Security & Architecture Audit Report

## Parking Management System v1

**Audit Date:** 2026-02-07  
**Auditor Role:** Expert Senior Software Architect & Lead Security Auditor  
**Audit Scope:** Entry Flow, Exit Flow, Activity Diagram Compliance

---

## Executive Summary

| Category                 | Status      | Critical Issues |
|--------------------------|-------------|-----------------|
| Architecture & Interface | ✅ PASS      | 0               |
| Entry Logic              | ⚠️ PARTIAL  | 3               |
| Exit Logic               | 🔴 CRITICAL | 5               |
| Fee Calculation          | ⚠️ PARTIAL  | 2               |
| Error Handling           | ⚠️ PARTIAL  | 2               |
| Race Conditions          | 🔴 CRITICAL | 2               |

**Overall Assessment:** ❌ **NOT Enterprise Ready** - Critical bugs must be fixed before production deployment.

---

## Phase 1: Architecture & Interface Consistency

### 1.1 In-File Interface Pattern ✅ PASS

All services and repositories correctly implement the In-File Interface pattern:

| Component                                                                  | Interface Defined | Concrete Struct Private          | Status |
|----------------------------------------------------------------------------|-------------------|----------------------------------|--------|
| [`ParkingTransactionService`](internal/services/parking_transaction.go:21) | ✅ Line 21-27      | ✅ `parkingTransactionService`    | PASS   |
| [`FeeCalculator`](internal/services/fee_calculator.go:24)                  | ✅ Line 24-27      | ✅ `feeCalculator`                | PASS   |
| [`ZoneService`](internal/services/zone.go:15)                              | ✅ Line 15-21      | ✅ `zoneService`                  | PASS   |
| [`ParkingTransactionRepository`](internal/repos/parking_transaction.go:8)  | ✅ Line 8-12       | ✅ `parkingTransactionRepository` | PASS   |
| [`ZoneRepository`](internal/repos/zone.go:8)                               | ✅ Line 8-9        | ✅ `zoneRepository`               | PASS   |

### 1.2 Container Wiring ✅ PASS

All dependencies in [`container.go`](internal/container/container.go:13) are declared as interfaces:

```go
// Line 59-61 - Correct interface types
ParkingTransactionRepo     repos.ParkingTransactionRepository
ParkingTransactionService  services.ParkingTransactionService
```

### 1.3 HashID Obfuscation ✅ PASS

- DTOs use `string` for hashed IDs: [
  `ParkingEntryRequest.VehicleTypeID`](internal/dto/requests/parking_transaction.go:6)
- Decoding happens in service layer before repository calls: [
  `parking_transaction.go:90-99`](internal/services/parking_transaction.go:90)
- Encoding happens in response mapping: [`parking_transaction.go:149-164`](internal/services/parking_transaction.go:149)

---

## Phase 2: Entry Logic Audit

### 2.1 RFID Check (Registered vs Guest Users) 🔴 MISSING

**Flow Requirement (Flow-Masuk-Parkir.txt:6-10):**

```
CheckDB{Check in Database: WHERE rfid_uid = ?}
CheckDB -->|Found| Registered[✅ REGISTERED USER]
CheckDB -->|Not Found| Guest[🆕 GUEST USER]
```

**Current Implementation:** ❌ **NOT IMPLEMENTED**

The [`RegisterEntry()`](internal/services/parking_transaction.go:73) method does NOT:

1. Check if RFID exists in the `customers` table
2. Distinguish between registered and guest users
3. Load vehicle data for registered users

**Impact:** System cannot differentiate user types, breaking the core business logic.

### 2.2 Double Entry Prevention ⚠️ PARTIAL

**Flow Requirement (Flow-Masuk-Parkir.txt:15-17):**

```
CheckActive{Check Active Parking: WHERE rfid_uid = ? AND exit_time IS NULL}
```

**Current Implementation:** Checks by `plate_number` instead of `rfid_uid`:

```go
// Line 105-107 - WRONG FIELD
Where("plate_number = ? AND status = ?", sanitizedPlate, "PARKED")
```

**Issue:** Should check by `rfid_uid` per the flow diagram.

### 2.3 Capacity Check 🔴 MISSING

**Flow Requirement (Flow-Masuk-Parkir.txt:54-58):**

```
CheckZoneFull{Zone Still Available?}
CheckZoneFull -->|Full| ZoneFull[❌ Zone Penuh!]
CheckZoneFull -->|Available| UpdateZone[...]
```

**Current Implementation:** ❌ **NOT IMPLEMENTED**

```go
// Line 117-123 - TODO comment, not implemented
// TODO: Implement zone occupancy check when ZoneService is available
```

**Impact:** System allows unlimited vehicles into a zone, causing physical overflow.

### 2.4 Atomic Transaction for Entry 🔴 MISSING

**Flow Requirement (Activity-Diagram.txt:111):**

```
BEGIN TRANSACTION
UPDATE zones SET occupied +1, available -1
UPDATE parking_transactions SET zone_status = occupied
COMMIT
```

**Current Implementation:** ❌ **NOT ATOMIC**

The [`RegisterEntry()`](internal/services/parking_transaction.go:144) method:

1. Creates transaction without DB transaction wrapper
2. Does NOT update zone occupancy
3. Does NOT create `TransactionZone` record

---

## Phase 3: Exit & Fee Logic Audit

### 3.1 Transaction Lookup ⚠️ PARTIAL

**Flow Requirement (Flow-Keluar-Parkir.txt:6):**

```
FindTX[Find Active Transaction: WHERE rfid_uid = ? AND exit_time IS NULL AND zone_status = occupied]
```

**Current Implementation:** Missing `zone_status` check:

```go
// Line 249-251 - Missing zone_status condition
Where("rfid_uid = ? AND exit_time IS NULL AND status = ?", req.RfidUID, "PARKED")
```

**Issue:** Should also check `zone_status = 'occupied'` in the `transaction_zones` table.

### 3.2 Fee Calculation 🔴 CRITICAL BUGS

**Flow Requirement (Activity-Diagram.txt:148-150):**

```
Calculate Duration: Zone Entry: 08:35, Now: 11:05, Duration: 2h 30m = 150 min
Calculate Fee: Zone: A Rate: Rp 2,000/h, Duration: 2.5 hours, Fee: 2.5 × 2000 = Rp 5,000
```

**Critical Bug #1: FreeMinutes NOT Handled**

The [`ZoneRate`](internal/models/zone_rate.go:14) model has `FreeMinutes` field, but [
`CalculateFeeWithDB()`](internal/services/fee_calculator.go:47) ignores it:

```go
// fee_calculator.go - FreeMinutes is NEVER used
// Should subtract free minutes before calculating fee
```

**Critical Bug #2: Duration Calculation Uses Entry Time, Not Zone Entry Time**

```go
// Line 272-278 - Uses transaction.EntryTime instead of transactionZone.EntryTime
feeResult, err := s.feeCalculator.CalculateFeeWithDB(
ctx,
transaction.EntryTime, // WRONG: Should be transactionZone.EntryTime
now,
...
)
```

**Flow Requirement:** Duration should be calculated from `zone_entry_time`, not `entry_time`.

### 3.3 Atomic Finalization 🔴 CRITICAL BUGS

**Flow Requirement (Activity-Diagram.txt:181-198):**

```
BEGIN TRANSACTION
  Release Zone Occupancy
  UPDATE zones SET occupied -1, available +1
  UPDATE parking_transactions SET zone_status = released
  UPDATE vehicles SET last_seen_at = NOW
  UPDATE customers SET total_visits +1
COMMIT TRANSACTION
```

**Critical Bug #3: Zone Occupancy Update is WRONG**

```go
// Line 320-324 - COMPLETELY WRONG LOGIC
newOccupiedCount := zone.MaximumCapacity - 1 // This makes no sense!
if newOccupiedCount < 0 {
newOccupiedCount = 0
}
newAvailableSlots := zone.MaximumCapacity - newOccupiedCount
```

**Expected Logic:**

```go
// Should decrement current occupied count, not use MaximumCapacity
newOccupiedCount := zone.OccupiedCount - 1
newAvailableSlots := zone.AvailableSlots + 1
```

**Critical Bug #4: Zone Model Missing Required Fields**

The [`Zone`](internal/models/zone.go:8) model is missing:

- `OccupiedCount int`
- `AvailableSlots int`
- `LastExitTime *time.Time`

These fields are referenced in the service but don't exist in the model!

**Critical Bug #5: No gorm.Expr for Atomic Updates**

```go
// Line 326-332 - Uses map update, NOT atomic
Updates(map[string]interface{}{
"occupied_count":  newOccupiedCount,
"available_slots": newAvailableSlots,
})
```

**Should use gorm.Expr for race-condition-safe updates:**

```go
Updates(map[string]interface{}{
"occupied_count":  gorm.Expr("occupied_count - ?", 1),
"available_slots": gorm.Expr("available_slots + ?", 1),
})
```

### 3.4 BeginTx Syntax Check ✅ PASS

No invalid `BeginTx` usage found. The code correctly uses `Begin()`:

```go
// Line 286 - Correct GORM syntax
tx := s.repo.GetDB().WithContext(ctx).Begin()
```

---

## Phase 4: Edge Cases & Error Handling

### 4.1 Race Conditions 🔴 CRITICAL

**Issue #1: Zone Occupancy Race Condition**

Multiple concurrent exit requests can cause negative occupancy:

```go
// Two requests read zone.OccupiedCount = 1 simultaneously
// Both calculate newOccupiedCount = 0
// Both update to 0, but should be -1 (invalid state)
```

**Fix:** Use `gorm.Expr` for atomic increment/decrement.

**Issue #2: Double Entry Race Condition**

The duplicate check at [`parking_transaction.go:105-115`](internal/services/parking_transaction.go:105) is not atomic:

```go
// Thread 1: Checks, finds no duplicate
// Thread 2: Checks, finds no duplicate
// Thread 1: Creates transaction
// Thread 2: Creates transaction (DUPLICATE!)
```

**Fix:** Use database-level unique constraint or `SELECT FOR UPDATE`.

### 4.2 Rollback Coverage ✅ PASS

All error paths in [`RegisterExit()`](internal/services/parking_transaction.go:237) have proper rollback:

| Line    | Operation             | Rollback          |
|---------|-----------------------|-------------------|
| 296-299 | Save transaction      | ✅ `tx.Rollback()` |
| 306-309 | Save transaction zone | ✅ `tx.Rollback()` |
| 314-317 | Fetch zone            | ✅ `tx.Rollback()` |
| 326-335 | Update zone           | ✅ `tx.Rollback()` |
| 340-345 | Update vehicle        | ✅ `tx.Rollback()` |
| 348-353 | Update customer       | ✅ `tx.Rollback()` |

### 4.3 DTO Validation Tags ⚠️ PARTIAL

| Field                                                                                 | Current                      | Expected                | Status                         |
|---------------------------------------------------------------------------------------|------------------------------|-------------------------|--------------------------------|
| [`ParkingEntryRequest.PlateNumber`](internal/dto/requests/parking_transaction.go:12)  | `required,max=20`            | `required,min=3,max=20` | ⚠️ Missing min                 |
| [`ParkingEntryRequest.RfidUID`](internal/dto/requests/parking_transaction.go:15)      | `omitempty,max=50`           | `required,max=50`       | ⚠️ Should be required per flow |
| [`ParkingExitRequest.PaymentMethod`](internal/dto/requests/parking_transaction.go:28) | `required,oneof=manual qris` | ✅ Correct               | PASS                           |

---

## Phase 5: Deliverables

### 5.1 Logical Mismatches Summary

| # | Flow Requirement               | Current Implementation  | Severity    |
|---|--------------------------------|-------------------------|-------------|
| 1 | RFID-based user type detection | Not implemented         | 🔴 Critical |
| 2 | Double entry check by RFID     | Checks by plate_number  | ⚠️ Medium   |
| 3 | Zone capacity validation       | Not implemented         | 🔴 Critical |
| 4 | Atomic entry transaction       | Not atomic              | 🔴 Critical |
| 5 | Duration from zone_entry_time  | Uses entry_time         | ⚠️ Medium   |
| 6 | FreeMinutes in fee calculation | Not implemented         | ⚠️ Medium   |
| 7 | Atomic zone occupancy update   | Non-atomic, wrong logic | 🔴 Critical |

### 5.2 Critical Bugs

1. **Zone Model Missing Fields** - `occupied_count`, `available_slots`, `last_exit_time` don't exist
2. **Zone Occupancy Logic is Completely Wrong** - Uses `MaximumCapacity - 1` instead of decrementing
3. **Race Condition in Zone Updates** - No atomic operations
4. **FreeMinutes Ignored** - Fee calculation doesn't honor free parking period
5. **Wrong Duration Base** - Should use zone entry time, not gate entry time

### 5.3 Actionable Code Snippets

#### Fix #1: Add Missing Zone Fields

```go
// internal/models/zone.go
type Zone struct {
// ... existing fields ...
OccupiedCount  int        `gorm:"default:0" json:"occupied_count"`
AvailableSlots int        `gorm:"default:0" json:"available_slots"`
LastExitTime   *time.Time `json:"last_exit_time"`
}
```

#### Fix #2: Correct Zone Occupancy Update with Atomic Operations

```go
// internal/services/parking_transaction.go - Replace lines 319-335
// Use gorm.Expr for atomic updates to prevent race conditions
if err := tx.Model(&models.Zone{}).
Where("id = ? AND occupied_count > 0", transactionZone.ActualZoneID).
Updates(map[string]interface{}{
"occupied_count":  gorm.Expr("occupied_count - ?", 1),
"available_slots": gorm.Expr("available_slots + ?", 1),
"last_exit_time":  now,
}).Error; err != nil {
tx.Rollback()
return responses.ParkingExitResponse{}, fmt.Errorf("failed to update zone occupancy: %w", err)
}
```

#### Fix #3: Implement FreeMinutes in Fee Calculator

```go
// internal/services/fee_calculator.go - Add after line 65
// Apply free minutes
effectiveDurationMinutes := durationMinutes - zoneRate.FreeMinutes
if effectiveDurationMinutes < 0 {
effectiveDurationMinutes = 0
}

// Recalculate hours with free minutes applied
effectiveDuration := time.Duration(effectiveDurationMinutes) * time.Minute
durationHours := math.Ceil(effectiveDuration.Hours())
if durationHours < 1 && effectiveDurationMinutes > 0 {
durationHours = 1 // Minimum 1 hour charge if any time beyond free period
}
```

#### Fix #4: Use Zone Entry Time for Duration Calculation

```go
// internal/services/parking_transaction.go - Replace line 272-278
// Use zone entry time for accurate duration calculation
feeResult, err := s.feeCalculator.CalculateFeeWithDB(
ctx,
transactionZone.EntryTime, // FIXED: Use zone entry time
now,
transactionZone.ActualZoneID,
transaction.VehicleTypeID,
s.repo.GetDB(),
)
```

#### Fix #5: Implement RFID-Based User Detection

```go
// internal/services/parking_transaction.go - Add before line 104
// Check if RFID is registered
var customer models.Customer
customerResult := s.repo.GetDB().WithContext(ctx).
Where("rfid_uid = ?", req.RfidUID).
First(&customer)

isRegistered := customerResult.Error == nil
var customerID *uint
if isRegistered {
customerID = &customer.ID
}

// Check for active parking by RFID (not plate number)
var existingTransaction models.ParkingTransaction
result := s.repo.GetDB().WithContext(ctx).
Where("rfid_uid = ? AND exit_time IS NULL", req.RfidUID).
First(&existingTransaction)

if result.Error == nil {
return responses.ParkingEntryResponse{}, errors.New("RFID already has active parking session")
}
```

#### Fix #6: Implement Zone Capacity Check

```go
// internal/services/parking_transaction.go - Add after line 99
// Validate zone capacity
var zone models.Zone
if err := s.repo.GetDB().WithContext(ctx).First(&zone, zoneID).Error; err != nil {
return responses.ParkingEntryResponse{}, fmt.Errorf("zone not found: %w", err)
}

if zone.OccupiedCount >= zone.MaximumCapacity {
return responses.ParkingEntryResponse{}, fmt.Errorf("zone %s is full (capacity: %d/%d)",
zone.Name, zone.OccupiedCount, zone.MaximumCapacity)
}
```

### 5.4 Enterprise Readiness Assessment

| Criteria                  | Status     | Notes                                       |
|---------------------------|------------|---------------------------------------------|
| Data Integrity            | ❌ FAIL     | Race conditions, wrong occupancy logic      |
| Business Logic Compliance | ❌ FAIL     | 7 major flow mismatches                     |
| Error Handling            | ⚠️ PARTIAL | Rollbacks OK, but missing validations       |
| Security                  | ⚠️ PARTIAL | HashID OK, but no rate limiting             |
| Scalability               | ❌ FAIL     | Non-atomic operations will fail under load  |
| Maintainability           | ✅ PASS     | Clean architecture, good interface patterns |

---

## Conclusion

**The system is NOT Enterprise Ready.**

### Immediate Actions Required (P0):

1. Add missing fields to Zone model
2. Fix zone occupancy update logic
3. Implement atomic operations with `gorm.Expr`
4. Implement RFID-based user detection
5. Implement zone capacity validation

### Short-term Actions (P1):

1. Fix duration calculation to use zone entry time
2. Implement FreeMinutes in fee calculation
3. Add database-level constraints for double entry prevention
4. Update DTO validation tags

### Recommended Testing:

1. Load testing with concurrent entry/exit requests
2. Edge case testing for zone capacity limits
3. Fee calculation verification against flow diagrams

---

*Report generated by Kilo Code Security Audit System*
