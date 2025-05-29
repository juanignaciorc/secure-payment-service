# Secure Payment Service

A microservice for handling bank account transfers with Go.

## Features

- Account management and balance tracking
- Transfer operations with pending status
- Transfer status updates via webhook
- Transactional database operations
- RESTful API endpoints

## Setup

1. Install Go 1.19+
2. Install PostgreSQL
3. Create the database:
   ```sql
   CREATE DATABASE payment_service;
   ```
4. Run the schema:
   ```sql
   \i internal/db/schema.sql
   ```

## Running the Service

```bash
# Build the service
go build -o server cmd/server/main.go

# Run the service
./server -port 8080
```

## API Endpoints

- `POST /api/transfers` - Create a new transfer
- `GET /api/transfers/info?id={id}` - Get transfer status
- `GET /api/transfers/status?id={id}&status={status}` - Update transfer status
- `GET /api/accounts/balance?account_id={id}` - Get account balance

## Example Usage

```bash
# Create a transfer
curl -X POST http://localhost:8080/api/transfers \
  -H "Content-Type: application/json" \
  -d '{"id": "transfer-123", "from_account": "acc-123", "to_account": "acc-456", "amount": 100.00}'

# Get transfer status
curl "http://localhost:8080/api/transfers/info?id=transfer-123"

# Update transfer status
curl "http://localhost:8080/api/transfers/status?id=transfer-123&status=COMPLETED"

# Get account balance
curl "http://localhost:8080/api/accounts/balance?account_id=acc-123"
```
