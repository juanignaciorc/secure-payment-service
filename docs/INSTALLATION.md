# Installation Guide

## Prerequisites

- Go 1.18 or higher
- PostgreSQL 13 or higher
- Docker (optional, for development)

## Installation Steps

### 1. Clone the Repository
```bash
git clone [repository-url]
cd secure-payment-service
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Set Up Database

#### Using Docker
```bash
docker run --name payment-service-db \
    -e POSTGRES_PASSWORD=postgres \
    -e POSTGRES_DB=payment_service \
    -p 5432:5432 \
    -d postgres
```

#### Using Local PostgreSQL
```bash
# Create database
createdb payment_service

# Create tables
psql payment_service < internal/migrations/0001_initial_schema.sql
psql payment_service < internal/migrations/0002_seed_data.sql
```

### 4. Configure Environment
Create a `.env` file with the following content:
```
DATABASE_DSN=postgresql://postgres:postgres@localhost:5432/payment_service?sslmode=disable
PORT=8080
JWT_SECRET=your-secret-key-here
```

### 5. Run the Application
```bash
# Start the main server
go run cmd/server/main.go

# Start the simulator (optional)
go run cmd/simulator/main.go
```

## Running Tests
```bash
go test ./...
```

## API Endpoints

### Transfer Operations
- POST `/api/transfers`: Create new transfer
- GET `/api/accounts/balance`: Get account balance
- POST `/api/webhook/transfer`: Process webhook notification

### Health Check
- GET `/health`: Check service status

## Development Environment

### Docker Compose
```bash
docker-compose up -d
```

### Local Development
```bash
# Run migrations
go run internal/migrations/apply_migrations.go

# Start server in development mode
go run -tags=dev cmd/server/main.go
```

## Troubleshooting

### Common Issues

1. Database Connection
   - Verify PostgreSQL is running
   - Check database credentials in .env
   - Ensure proper network access

2. Missing Dependencies
   - Run `go mod tidy`
   - Verify Go version

3. Migration Errors
   - Drop and recreate database
   - Run migrations in order