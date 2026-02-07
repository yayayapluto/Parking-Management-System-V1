# Parking Exit API - Usage Examples

## Overview

This document provides practical examples for using the Parking Exit API endpoint.

## Base URL

```
http://localhost:8080/api/v1
```

## Endpoint

```
POST /transactions/exit
```

## Request Headers

```
Content-Type: application/json
Authorization: Bearer <JWT_TOKEN> (if required)
```

## Examples

### Example 1: Successful Exit with Manual Payment

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID123456789",
    "payment_method": "manual"
  }'
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Vehicle exit recorded successfully",
  "data": {
    "transaction_id": "eJ5xK9mL2pQ",
    "rfid_uid": "RFID123456789",
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

**Explanation:**

- Vehicle parked for 150 minutes (2 hours 30 minutes)
- Rounded up to 3 hours for billing
- Hourly rate: 10,000 IDR
- Total fee: 3 hours × 10,000 = 30,000 IDR

---

### Example 2: Exit with QRIS Payment

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID987654321",
    "payment_method": "qris"
  }'
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Vehicle exit recorded successfully",
  "data": {
    "transaction_id": "aB3cD4eF5gH",
    "rfid_uid": "RFID987654321",
    "plate_number": "D 5678 XYZ",
    "entry_time": "2026-02-07T08:00:00Z",
    "exit_time": "2026-02-07T14:00:00Z",
    "total_duration": 360,
    "total_fee": 60000.00,
    "payment_status": "paid",
    "payment_method": "qris"
  }
}
```

**Explanation:**

- Vehicle parked for 360 minutes (6 hours)
- Hourly rate: 10,000 IDR
- Total fee: 6 hours × 10,000 = 60,000 IDR
- Payment method: QRIS (QR code payment)

---

### Example 3: Short Duration Parking (Less than 1 Hour)

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID111222333",
    "payment_method": "manual"
  }'
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Vehicle exit recorded successfully",
  "data": {
    "transaction_id": "xY9zW8vU7tS",
    "rfid_uid": "RFID111222333",
    "plate_number": "A 9999 DEF",
    "entry_time": "2026-02-07T15:45:00Z",
    "exit_time": "2026-02-07T16:15:00Z",
    "total_duration": 30,
    "total_fee": 10000.00,
    "payment_status": "paid",
    "payment_method": "manual"
  }
}
```

**Explanation:**

- Vehicle parked for 30 minutes
- Minimum charge applies: 1 hour
- Hourly rate: 10,000 IDR
- Total fee: 1 hour × 10,000 = 10,000 IDR

---

### Example 4: No Active Parking Found

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "NONEXISTENT_RFID",
    "payment_method": "manual"
  }'
```

**Response (404 Not Found):**

```json
{
  "success": false,
  "message": "no active parking found",
  "error": "no active parking found for the given RFID"
}
```

**Explanation:**

- RFID tag not found in active parking transactions
- Vehicle may not have entered the parking system
- Vehicle may have already exited

---

### Example 5: Invalid Payment Method

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID123456789",
    "payment_method": "credit_card"
  }'
```

**Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "invalid request payload",
  "error": "Key: 'ParkingExitRequest.PaymentMethod' Error:Field validation for 'PaymentMethod' failed on the 'oneof' tag"
}
```

**Explanation:**

- Payment method must be either "manual" or "qris"
- Invalid payment methods are rejected at validation stage

---

### Example 6: Missing Required Field

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID123456789"
  }'
```

**Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "invalid request payload",
  "error": "Key: 'ParkingExitRequest.PaymentMethod' Error:Field validation for 'PaymentMethod' failed on the 'required' tag"
}
```

**Explanation:**

- Payment method is required
- Request must include both rfid_uid and payment_method

---

### Example 7: Long Duration Parking with Daily Max Rate

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{
    "rfid_uid": "RFID444555666",
    "payment_method": "manual"
  }'
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Vehicle exit recorded successfully",
  "data": {
    "transaction_id": "pQ1rS2tU3vW",
    "rfid_uid": "RFID444555666",
    "plate_number": "C 3333 GHI",
    "entry_time": "2026-02-06T10:00:00Z",
    "exit_time": "2026-02-07T18:00:00Z",
    "total_duration": 1440,
    "total_fee": 100000.00,
    "payment_status": "paid",
    "payment_method": "manual"
  }
}
```

**Explanation:**

- Vehicle parked for 1440 minutes (24 hours)
- Calculated fee: 24 hours × 10,000 = 240,000 IDR
- Daily max rate applied: 100,000 IDR
- Total fee: 100,000 IDR (capped at daily maximum)

---

## Response Status Codes

| Status Code | Meaning               | Scenario                                    |
|-------------|-----------------------|---------------------------------------------|
| 200         | OK                    | Successful exit recorded                    |
| 400         | Bad Request           | Invalid request payload or validation error |
| 404         | Not Found             | No active parking found for RFID            |
| 500         | Internal Server Error | Database or system error                    |

## Receipt Information

The response data can be used to generate a parking receipt with the following information:

```
═══════════════════════════════════════
        PARKING RECEIPT
═══════════════════════════════════════
Transaction ID: eJ5xK9mL2pQ
Plate Number:   B 1234 ABC
RFID UID:       RFID123456789

Entry Time:     2026-02-07 10:00:00
Exit Time:      2026-02-07 12:30:00
Duration:       2 hours 30 minutes

Hourly Rate:    10,000 IDR
Duration Hours: 3 hours
Total Fee:      30,000 IDR

Payment Method: Manual
Payment Status: Paid

═══════════════════════════════════════
Thank you for using our parking service!
═══════════════════════════════════════
```

## Integration Notes

1. **RFID Reading**: Ensure RFID reader is properly configured to capture the UID
2. **Timestamp Accuracy**: System uses server time (UTC) for all calculations
3. **Concurrency**: Multiple simultaneous exits are handled safely with database transactions
4. **Idempotency**: Calling exit twice with same RFID will fail on second attempt (transaction already exited)
5. **Rate Lookup**: Rates are fetched from zone_rates table based on zone, vehicle type, and date

## Error Handling

Always check the `success` field in the response:

- `success: true` → Transaction completed successfully
- `success: false` → Error occurred, check `error` field for details

## Testing with cURL

For testing without authentication:

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -d '{"rfid_uid":"TEST123","payment_method":"manual"}'
```

For testing with authentication:

```bash
curl -X POST http://localhost:8080/api/v1/transactions/exit \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"rfid_uid":"TEST123","payment_method":"manual"}'
```
