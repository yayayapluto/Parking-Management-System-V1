# 🏗️ Complete Project Structure (Directories & Files Only)

## 📁 Full Directory & File Structure

```
parking-system/
│
├── cmd/
│   └── main.go
│
├── internal/
│   │
│   ├── models/
│   │   ├── vehicle_type.go           // GORM model
│   │   ├── payment_method.go         // GORM model
│   │   ├── holiday.go                // GORM model
│   │   ├── zone_type.go              // GORM model
│   │   ├── zone.go                   // GORM model
│   │   ├── zone_rate.go              // GORM model
│   │   ├── permission.go             // GORM model
│   │   ├── role.go                   // GORM model
│   │   ├── user.go                   // GORM model
│   │   ├── customer_regist_source.go // GORM model
│   │   ├── customer.go               // GORM model
│   │   ├── vehicle.go                // GORM model
│   │   ├── parking_transaction.go    // GORM model
│   │   ├── transaction_zone.go       // GORM model
│   │   ├── transaction_ocr_data.go   // GORM model
│   │   ├── transaction_event.go      // GORM model
│   │   ├── payment.go                // GORM model
│   │   ├── refund.go                 // GORM model
│   │   ├── lost_ticket_fee.go        // GORM model
│   │   ├── manual_correction.go      // GORM model
│   │   ├── zone_occupancy_log.go     // GORM model
│   │   ├── ocr_log.go                // GORM model
│   │   ├── user_activity_log.go      // GORM model
│   │   ├── system_log.go             // GORM model
│   │   ├── shift_report.go           // GORM model
│   │   ├── daily_settlement.go       // GORM model
│   │   └── report_cache.go           // GORM model
│   │
│   ├── repositories/
│   │   ├── vehicle_repository.go
│   │   ├── zone_repository.go
│   │   ├── transaction_repository.go
│   │   ├── payment_repository.go
│   │   ├── ocr_repository.go
│   │   ├── operator_repository.go
│   │   ├── zone_rate_repository.go
│   │   └── zone_occupancy_repository.go
│   │
│   ├── services/
│   │   ├── vehicle_service.go
│   │   ├── zone_service.go
│   │   ├── entry_service.go
│   │   ├── zone_entry_service.go
│   │   ├── exit_service.go
│   │   ├── payment_service.go
│   │   ├── ocr_service.go
│   │   ├── fee_calculator.go
│   │   ├── receipt_generator.go
│   │   ├── report_service.go
│   │   └── security_service.go
│   │
│   ├── handlers/
│   │   ├── vehicle_handler.go
│   │   ├── zone_handler.go
│   │   ├── entry_handler.go
│   │   ├── exit_handler.go
│   │   ├── payment_handler.go
│   │   ├── webhook_handler.go
│   │   ├── report_handler.go
│   │   └── operator_handler.go
│   │
│   ├── dto/
│   │   ├── requests/
│   │   │   ├── entry_request.go
│   │   │   ├── zone_select_request.go
│   │   │   ├── zone_tap_request.go
│   │   │   ├── exit_request.go
│   │   │   ├── exit_confirm_request.go
│   │   │   ├── payment_request.go
│   │   │   ├── vehicle_request.go
│   │   │   ├── zone_request.go
│   │   │   ├── operator_request.go
│   │   │   └── login_request.go
│   │   │
│   │   └── responses/
│   │       ├── entry_response.go
│   │       ├── zone_availability_response.go
│   │       ├── exit_response.go
│   │       ├── payment_response.go
│   │       ├── zone_response.go
│   │       ├── vehicle_response.go
│   │       ├── transaction_response.go
│   │       ├── report_response.go
│   │       ├── common_response.go
│   │       └── error_response.go
│   │
│   ├── clients/
│   │   ├── ocr_client.go
│   │   ├── midtrans_client.go
│   │   ├── storage_client.go
│   │   └── camera_client.go
│   │
│   ├── circuitbreaker/
│   │   ├── breaker.go
│   │   └── config.go
│   │
│   ├── middlewares/
│   │   ├── auth.go
│   │   ├── logger.go
│   │   ├── cors.go
│   │   ├── error_handler.go
│   │   ├── rate_limiter.go
│   │   └── request_id.go
│   │
│   ├── routes/
│   │   └── routes.go
│   │
│   └── app/
│       ├── container.go
│       ├── app.go
│       ├── database.go
│       └── server.go
│
├── pkg/
│   ├── config/
│   │   ├── config.go
│   │   └── env.go
│   │
│   ├── logger/
│   │   ├── logger.go
│   │   └── zap.go
│   │
│   ├── errors/
│   │   ├── errors.go
│   │   ├── codes.go
│   │   └── handler.go
│   │
│   ├── validator/
│   │   ├── validator.go
│   │   └── rules.go
│   │
│   ├── database/
│   │   ├── postgres.go
│   │   └── transaction.go
│   │
│   └── helpers/
│       ├── string.go
│       ├── time.go
│       ├── response.go
│       ├── pagination.go
│       └── hash.go
│
├── config/
│   ├── config.yaml
│   ├── config.dev.yaml
│   ├── config.prod.yaml
│   └── config.example.yaml
│
├── scripts/
│   ├── build.sh
│   ├── deploy.sh
│   ├── migrate-up.sh
│   ├── migrate-down.sh
│   └── seed.sh
│
├── tests/
│   ├── unit/
│   │   ├── services/
│   │   │   ├── entry_service_test.go
│   │   │   ├── exit_service_test.go
│   │   │   ├── payment_service_test.go
│   │   │   └── fee_calculator_test.go
│   │   │
│   │   ├── repositories/
│   │   │   ├── vehicle_repository_test.go
│   │   │   ├── zone_repository_test.go
│   │   │   └── transaction_repository_test.go
│   │   │
│   │   └── helpers/
│   │       ├── string_test.go
│   │       └── time_test.go
│   │
│   ├── integration/
│   │   ├── entry_flow_test.go
│   │   ├── exit_flow_test.go
│   │   ├── payment_flow_test.go
│   │   └── zone_management_test.go
│   │
│   ├── e2e/
│   │   ├── complete_parking_flow_test.go
│   │   └── concurrent_parking_test.go
│   │
│   ├── mocks/
│   │   ├── mock_vehicle_repository.go
│   │   ├── mock_zone_repository.go
│   │   ├── mock_transaction_repository.go
│   │   ├── mock_ocr_client.go
│   │   └── mock_midtrans_client.go
│   │
│   └── fixtures/
│       ├── vehicles.json
│       ├── zones.json
│       ├── transactions.json
│       └── test_images/
│           ├── plate_sample_1.jpg
│           ├── plate_sample_2.jpg
│           └── plate_sample_3.jpg
│
├── docs/
│   ├── api/
│   │   ├── openapi.yaml
│   │   ├── postman_collection.json
│   │   └── api_documentation.md
│   │
│   ├── architecture/
│   │   ├── system_design.md
│   │   ├── database_schema.md
│   │   ├── flow_diagrams.md
│   │   └── deployment.md
│   │
│   └── guides/
│       ├── setup_guide.md
│       ├── developer_guide.md
│       ├── deployment_guide.md
│       └── troubleshooting.md
│
├── deployments/
│   ├── docker/
│   │   ├── Dockerfile
│   │   ├── Dockerfile.dev
│   │   └── .dockerignore
│   │
│   ├── kubernetes/
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   ├── configmap.yaml
│   │   ├── secret.yaml
│   │   └── ingress.yaml
│   │
│   └── railway/
│       ├── railway.toml
│       └── railway.json
│
├── storage/
│   ├── photos/
│   │   ├── entry/
│   │   │   └── .gitkeep
│   │   ├── exit/
│   │   │   └── .gitkeep
│   │   └── plates/
│   │       └── .gitkeep
│   │
│   └── receipts/
│       └── .gitkeep
│
├── .env
├── .env.example
├── .env.dev
├── .env.prod
├── .gitignore
├── .editorconfig
├── .golangci.yaml
├── go.mod
├── go.sum
├── Makefile
├── docker-compose.yml
├── docker-compose.dev.yml
├── docker-compose.prod.yml
├── README.md
├── CHANGELOG.md
├── CONTRIBUTING.md
└── LICENSE

```

---

## 📊 Python OCR Service Structure

```
ocr-service/
│
├── app/
│   ├── __init__.py
│   ├── main.py
│   │
│   ├── routers/
│   │   ├── __init__.py
│   │   ├── ocr.py
│   │   └── health.py
│   │
│   ├── services/
│   │   ├── __init__.py
│   │   ├── preprocessor.py
│   │   ├── easy_ocr_engine.py
│   │   ├── paddle_ocr_engine.py
│   │   ├── tesseract_engine.py
│   │   ├── ensemble.py
│   │   └── ocr_service.py
│   │
│   ├── validators/
│   │   ├── __init__.py
│   │   ├── plate_validator.py
│   │   └── image_validator.py
│   │
│   ├── models/
│   │   ├── __init__.py
│   │   ├── request.py
│   │   └── response.py
│   │
│   ├── utils/
│   │   ├── __init__.py
│   │   ├── logger.py
│   │   ├── image_utils.py
│   │   └── config.py
│   │
│   └── config/
│       ├── __init__.py
│       ├── settings.py
│       └── constants.py
│
├── models/
│   ├── easyocr/
│   │   └── .gitkeep
│   ├── paddleocr/
│   │   └── .gitkeep
│   └── tesseract/
│       └── .gitkeep
│
├── tests/
│   ├── __init__.py
│   ├── test_preprocessor.py
│   ├── test_ocr_engines.py
│   ├── test_validator.py
│   └── test_api.py
│
├── sample_images/
│   ├── plate_1.jpg
│   ├── plate_2.jpg
│   └── plate_3.jpg
│
├── .env
├── .env.example
├── .gitignore
├── requirements.txt
├── requirements-dev.txt
├── Dockerfile
├── docker-compose.yml
├── pytest.ini
├── README.md
└── run.py

```

---

## 📋 Configuration Files

### **Golang - .gitignore**

```
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# Environment variables
.env
.env.local
.env.*.local

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Logs
*.log
logs/

# Storage
storage/photos/*
!storage/photos/.gitkeep
storage/receipts/*
!storage/receipts/.gitkeep

# Temporary files
tmp/
temp/

```

---

### **Golang - .env.example**

```
# Application
APP_NAME=parking-system
APP_ENV=development
APP_PORT=8080
APP_DEBUG=true

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=parking_db
DB_SSLMODE=disable
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5

# OCR Service
OCR_SERVICE_URL=http://localhost:8000
OCR_TIMEOUT=30s

# Midtrans
MIDTRANS_SERVER_KEY=your-server-key
MIDTRANS_CLIENT_KEY=your-client-key
MIDTRANS_ENVIRONMENT=sandbox

# Railway Storage
STORAGE_BUCKET_URL=https://your-bucket.railway.app
STORAGE_ACCESS_KEY=your-access-key
STORAGE_SECRET_KEY=your-secret-key

# Camera
CAMERA_API_URL=http://192.168.1.100:8080
CAMERA_USERNAME=admin
CAMERA_PASSWORD=admin123

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Circuit Breaker
CB_MAX_FAILURES=5
CB_TIMEOUT=10s
CB_RESET_TIMEOUT=30s

# Logging
LOG_LEVEL=debug
LOG_FORMAT=json

```

---

### **Golang - Makefile**

```makefile
.PHONY: help build run test clean migrate-up migrate-down docker-build docker-run

help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make migrate-up   - Run database migrations up"
	@echo "  make migrate-down - Run database migrations down"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"

build:
	go build -o bin/parking-system cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -rf tmp/
	go clean

migrate-up:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/parking_db?sslmode=disable" up

migrate-down:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/parking_db?sslmode=disable" down

docker-build:
	docker build -t parking-system:latest -f deployments/docker/Dockerfile .

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

lint:
	golangci-lint run

fmt:
	go fmt ./...

tidy:
	go mod tidy

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

```

---

### **Golang - docker-compose.yml**

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: parking-postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: parking_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - parking-network

  parking-api:
    build:
      context: .
      dockerfile: deployments/docker/Dockerfile
    container_name: parking-api
    environment:
      - APP_ENV=development
      - DB_HOST=postgres
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    networks:
      - parking-network

  ocr-service:
    build:
      context: ../ocr-service
      dockerfile: Dockerfile
    container_name: parking-ocr
    ports:
      - "8000:8000"
    networks:
      - parking-network

volumes:
  postgres_data:

networks:
  parking-network:
    driver: bridge

```

---

### **Python - .gitignore**

```
# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
*.egg-info/
.installed.cfg
*.egg

# PyInstaller
*.manifest
*.spec

# Unit test / coverage reports
htmlcov/
.tox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
.hypothesis/
.pytest_cache/

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Models
models/easyocr/*
!models/easyocr/.gitkeep
models/paddleocr/*
!models/paddleocr/.gitkeep

# Logs
*.log
logs/

```

---

### **Python - requirements.txt**

```
fastapi==0.104.1
uvicorn[standard]==0.24.0
pydantic==2.5.0
pydantic-settings==2.1.0
python-multipart==0.0.6
opencv-python==4.8.1.78
numpy==1.24.3
Pillow==10.1.0
easyocr==1.7.1
paddleocr==2.7.0.3
paddlepaddle==2.5.2
python-dotenv==1.0.0
requests==2.31.0

```

---

### **Python - requirements-dev.txt**

```
-r requirements.txt
pytest==7.4.3
pytest-cov==4.1.0
pytest-asyncio==0.21.1
black==23.12.0
flake8==6.1.0
mypy==1.7.1
httpx==0.25.2

```

---

## 📝 README Structure

### **README.md Sections**

```
# Parking System

## Features
## Tech Stack
## Architecture
## Prerequisites
## Installation
  - Clone Repository
  - Setup Database
  - Environment Variables
  - Run Migrations
  - Run Application
## API Documentation
## Development
  - Running Tests
  - Code Quality
  - Database Migrations
## Deployment
  - Docker
  - Railway
  - Kubernetes
## Project Structure
## Contributing
## License

```

---