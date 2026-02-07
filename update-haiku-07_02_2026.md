# Parking Management System - Code Logic & Flow Explanation

## Architecture Overview

This is a **Go-based REST API** using the **Fiber framework** with a **layered architecture** (Models → DTOs → Services → Handlers). The system manages parking zones, rates, users, roles, and authentication.

---

## 1. **Zone Management** (`zone.go` files)

### Data Flow:
```
Client Request → Handler → Service → Repository → Database
                                ↓
                         ID Obfuscation
                         Validation
                         Mapping
```

### [`internal/models/zone.go`](internal/models/zone.go:1)
- **Zone struct**: Database model with fields like `Name`, `Location`, `MaximumCapacity`, `DefaultFee`
- **Relations**: References `ZoneType` via foreign key
- **Nullable fields**: `OperatingHourStart/End` use `sql.NullString` for optional times

### [`internal/dto/requests/zone.go`](internal/dto/requests/zone.go:1)
- **CreateZoneRequest**: All required fields (name, location, capacity, fee)
- **UpdateZoneRequest**: All fields optional (pointers for numeric values to distinguish null from zero)
- **Validation tags**: `required`, `max=100`, `min=1` using go-playground/validator

### [`internal/dto/responses/zone.go`](internal/dto/responses/zone.go:1)
- Converts database model to JSON response
- All IDs are hashed strings (obfuscated) for security
- Timestamps formatted as strings

### [`internal/services/zone.go`](internal/services/zone.go:1)
**Key Methods:**

1. **GetAll()** (lines 35-53)
    - Calls repo with pagination + searchable columns ("name", "location")
    - Maps all results to response DTOs
    - Returns paginated response with metadata

2. **GetByID()** (lines 55-67)
    - Decodes hashed ID using obfuscator
    - Fetches from repo
    - Maps to response

3. **Create()** (lines 69-107)
    - Validates request struct
    - Decodes ZoneTypeID from hash
    - Converts optional time strings to `sql.NullString`
    - Creates model with `IsActive=true`, `IsInMaintenance=false`
    - Saves to database

4. **Update()** (lines 109-170)
    - **Patch pattern**: Only updates fields that are provided (non-empty/non-nil)
    - Fetches existing record
    - Conditionally updates each field
    - Saves changes

5. **Delete()** (lines 172-179)
    - Decodes ID and calls repo delete

6. **mapToResponse()** (lines 183-215)
    - Encodes IDs back to hashes
    - Handles nullable fields (checks `.Valid` flag)
    - Formats timestamps

### [`internal/handlers/zone.go`](internal/handlers/zone.go:1)
- **HTTP endpoints** with Swagger documentation
- **GetAll**: `GET /zones` with pagination params
- **Create**: `POST /zones` with JSON body
- **GetByID**: `GET /zones/{id}`
- **Update**: `PUT /zones/{id}`
- **Delete**: `DELETE /zones/{id}`
- All handlers parse requests, call service, return responses via `BaseHandler` methods

---

## 2. **Zone Rate Management** (Similar pattern to Zone)

### [`internal/models/zone_rate.go`](internal/models/zone_rate.go:1)
- Pricing rules per zone + vehicle type
- Fields: `HourlyRate`, `DailyMaxRate`, `FreeMinutes`
- Supports weekend/holiday rates
- Date range: `ValidFrom` to `ValidTo`
- Optional effective hours: `EffectiveHourFrom/To`

### [`internal/dto/requests/zone_rate.go`](internal/dto/requests/zone_rate.go:1)
- **CreateZoneRateRequest**: Zone ID, Vehicle Type ID, rates, dates
- **UpdateZoneRateRequest**: All fields optional

### [`internal/services/zone_rate.go`](internal/services/zone_rate.go:1)
**Key Logic:**

1. **Create()** (lines 70-136)
    - Decodes ZoneID and VehicleTypeID from hashes
    - **Date parsing**: Converts "2006-01-02" format strings to `time.Time`
    - Handles optional HolidayID (nullable pointer)
    - Converts optional time strings to `sql.NullString`
    - Creates with `IsActive=true`

2. **Update()** (lines 138-237)
    - Patch pattern: Only updates provided fields
    - **Date parsing**: Re-parses dates if provided
    - Handles nullable HolidayID (can set to nil if empty string)

3. **Delete()** (lines 239-246)
    - Standard delete by decoded ID

---

## 3. **User Management**

### [`internal/dto/requests/user.go`](internal/dto/requests/user.go:1)
- **CreateUserRequest**: Username (alphanum), email, password (min 8 chars), role
- **UpdateUserRequest**: Optional fields only
- **ChangePasswordRequest**: Old + new password validation

### [`internal/services/user.go`](internal/services/user.go:1)
**Key Methods:**

1. **Create()** (lines 71-105)
    - Validates request
    - Decodes RoleID from hash
    - **Password hashing**: Uses `bcrypt.GenerateFromPassword()` with default cost
    - Creates user with `IsActive=true`, `IsLocked=false`, `FailedLoginAttempts=0`

2. **Update()** (lines 108-151)
    - Patch pattern: Only updates non-empty fields
    - Decodes RoleID if provided

3. **GetByEmail()** & **GetByUsername()** (lines 164-193)
    - Fetches all users and searches in-memory
    - Returns error if not found

4. **mapToResponse()** (lines 196-215)
    - Encodes user ID and role ID to hashes
    - Handles nullable `LastLoginAt` timestamp

### [`internal/handlers/user_handler.go`](internal/handlers/user_handler.go:1)
- Standard CRUD endpoints: GET all, GET by ID, POST create, PUT update, DELETE
- All use pagination for list endpoints

---

## 4. **Role Management**

### [`internal/services/role.go`](internal/services/role.go:1)
- Similar CRUD pattern to Zone/User
- **Create()**: Decodes PermissionID, creates role with `CreatedBy` field
- **Update()**: Patch pattern for optional fields
- Searchable by "name"

---

## 5. **Authentication & Authorization**

### [`internal/services/auth.go`](internal/services/auth.go:1)
**Key Methods:**

1. **Login()** (lines 32-61)
    - Finds user by email
    - Checks if user is active
    - Checks if user is locked (with `LockedUntil` timestamp)
    - **Password verification**: Uses `bcrypt.CompareHashAndPassword()`
    - Returns user response

2. **LoginWithUsername()** (lines 64-93)
    - Same logic but searches by username instead

3. **VerifyPassword()** (lines 96-99)
    - Wrapper around bcrypt comparison
    - Returns boolean

4. **GenerateToken()** (lines 102-123)
    - Creates JWT with claims: `user_id`, `username`, `email`, `role_id`, `role_name`, `permissions`
    - Sets expiration based on config (`JWTConfig.ExpiryHours`)
    - Signs with HS256 using secret key
    - Returns token string + expiration time

### [`internal/handlers/auth_handler.go`](internal/handlers/auth_handler.go:1)
- **POST /auth/login**: Email + password login
- **POST /auth/login-username**: Username + password login
- Both return user info (token generation happens elsewhere)

### [`internal/middlewares/auth.go`](internal/middlewares/auth.go:1)
**Two Middleware Functions:**

1. **JWTMiddleware()** (lines 12-69)
    - **Mandatory authentication**
    - Extracts "Bearer <token>" from Authorization header
    - Parses JWT with HS256 verification
    - Validates token signature and expiration
    - Stores claims in context locals: `user_id`, `username`, `email`, `role_id`, `role_name`, `permissions`
    - Returns 401 if missing/invalid

2. **OptionalJWTMiddleware()** (lines 72-110)
    - **Optional authentication**
    - Same logic but doesn't fail if token missing
    - Continues to next handler regardless

---

## 6. **Key Design Patterns**

### **ID Obfuscation**
- All IDs are hashed in requests/responses using `IDObfuscator`
- Database stores actual uint IDs
- Prevents ID enumeration attacks

### **Patch Updates**
- Update requests use optional fields (pointers for numbers)
- Only provided fields are updated
- Existing values preserved if not specified

### **Nullable Fields**
- Optional times/dates use `sql.NullString`
- Checked with `.Valid` flag before use
- Allows distinguishing between null and empty string

### **Validation**
- Uses `go-playground/validator/v10`
- Struct tags: `required`, `email`, `max=100`, `min=0`, `alphanum`
- Validated in service layer before database operations

### **Error Handling**
- Services return errors directly
- Handlers propagate errors (likely caught by global error handler)
- No explicit HTTP status codes in services

### **Pagination**
- `PaginationRequest` with page/limit/search
- Repos return (models, total count, error)
- Services create `PageResponse` with metadata

---

## 7. **Request/Response Flow Example**

**Create Zone Request:**
```
POST /zones
{
  "name": "Zone A",
  "zone_type_id": "abc123xyz",  // hashed
  "location": "Building 1",
  "maximum_capacity": 50,
  "default_fee": 10000.00,
  "is_24_hours": true
}
    ↓
ZoneHandler.Create()
    ↓
ZoneService.Create()
  - Validate request
  - Decode zone_type_id hash → uint
  - Create Zone model
  - Save to database
  - Map to response
    ↓
ZoneResponse
{
  "id": "xyz789abc",  // hashed
  "name": "Zone A",
  "zone_type_id": "abc123xyz",
  "location": "Building 1",
  "maximum_capacity": 50,
  "default_fee": 10000.00,
  "is_24_hours": true,
  "created_at": "2026-02-07T05:35:00Z"
}
```

---

## 8. **Security Features**

- **Password hashing**: bcrypt with default cost
- **JWT tokens**: HS256 signed with secret key
- **ID obfuscation**: Hashed IDs prevent enumeration
- **Account locking**: Tracks failed attempts and lock time
- **Authorization middleware**: Validates tokens and extracts claims
- **Input validation**: Struct-level validation before processing

This architecture ensures clean separation of concerns, security, and maintainability.