# Secure Payment Service

A microservice for handling secure bank account transfers with Go. This service includes three main components:
- Payment service (main API)
- Transfer processor (simulates external bank processing)
- PostgreSQL database

## 🚀 Quick Start

1. **Prerequisites**
   - Docker and Docker Compose installed

2. **Build and Run Services**
   ```bash
   # 1. Clean up any existing containers and networks
   docker compose down -v
   
   # 2. Build and start all services
   docker compose up -d --build
   
   # Verify services are running
   docker ps
   ```

3. **Service URLs**
   - Payment service: `http://localhost:8080`
   - Transfer processor: `http://localhost:8081`
   - Database: `postgres://postgres:postgres@localhost:5432/payment_db`

## 🧪 Complete Test Flow

### 1. Initial Setup
- All services will start automatically with Docker
- Database is initialized with seed data (two accounts: acc-100 and acc-200)
- The transfer processor simulates external bank processing with 80% success rate

### 2. Test Accounts (Pre-configured)
Two accounts are pre-configured in the database:
- `acc-100`: Initial balance 1000.00
- `acc-200`: Initial balance 500.00

You can verify their balances:
```bash
curl "http://localhost:8080/api/accounts/balance?account_id=acc-100"
curl "http://localhost:8080/api/accounts/balance?account_id=acc-200"
```

### 3. Create a Transfer
```bash
# Create a transfer request
curl -X POST http://localhost:8080/api/transfers \
  -H "Content-Type: application/json" \
  -d '{"id": "transfer-001", "from_account": "acc-100", "to_account": "acc-200", "amount": 150.00}'
```

### 4. Check Transfer Status
```bash
# Check initial status (will be PENDING)
curl "http://localhost:8080/api/transfers/info?id=transfer-001"

# Wait 1-5 seconds for processing
sleep 5

# Check final status (COMPLETED or FAILED)
curl "http://localhost:8080/api/transfers/info?id=transfer-001"
```

### 5. Verify Balances
```bash
# Check sender's balance (should be reduced by 150 if successful)
curl "http://localhost:8080/api/accounts/balance?account_id=acc-100"

# Check receiver's balance (should be increased by 150 if successful)
curl "http://localhost:8080/api/accounts/balance?account_id=acc-200"
```

## 📝 API Documentation

### Health Check
- `GET /health`
  - Check service health
  - Response: `"OK"`

### Accounts API
- `POST /api/accounts`
  - Create new account
  - Request: `{
      "id": "account-id",
      "balance": 1000.00
    }`
  - Response: `{
      "id": "account-id",
      "balance": 1000.00,
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }`

- `GET /api/accounts/balance?account_id=<id>`
  - Get account balance
  - Response: `{
      "balance": 1000.00
    }`

### Transfers API
- `POST /api/transfers`
  - Create new transfer
  - Request: `{
      "id": "transfer-id",
      "from_account": "from-id",
      "to_account": "to-id",
      "amount": 100.00
    }`
  - Response: `{
      "id": "transfer-id",
      "from_account": "from-id",
      "to_account": "to-id",
      "amount": 100.00,
      "status": "PENDING",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }`

- `GET /api/transfers/info?id=<transfer-id>`
  - Get transfer status
  - Response: `{
      "id": "transfer-id",
      "from_account": "from-id",
      "to_account": "to-id",
      "amount": 100.00,
      "status": "PENDING/COMPLETED/FAILED",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }`

### Webhook API
- `POST /api/webhook/transfer`
  - Handle transfer status updates from the transfer processor
  - Request: `{
      "transfer_id": "transfer-id",
      "status": "COMPLETED/FAILED",
      "message": "optional message"
    }`
  - Response: HTTP 200 if successful, error otherwise

### Transfer Processor API
- `POST /process_transfer`
  - Internal endpoint for transfer processor
  - Request: `{
      "transfer_id": "transfer-id",
      "amount": 100.00
    }`
  - Response: HTTP 202 if accepted, error otherwise

## 🛠️ Troubleshooting

1. **If services don't start**
   - Ensure Docker and Docker Compose are installed
   - Run `docker compose down -v` to clean up previous containers and volumes
   - Try rebuilding with `docker compose up -d --build`

2. **If database connection fails**
   - Wait 10-15 seconds for database initialization
   - Check if PostgreSQL is running with `docker ps`
   - Verify database logs with `docker logs secure-payment-service-postgres-1`

3. **If transfer fails**
   - Verify account balances are sufficient
   - Check transfer ID format (should be unique)
   - Ensure accounts exist
   - Wait for processing time (1-5 seconds)
   - Check transfer processor logs with `docker logs secure-payment-service-transfer-processor-1`

4. **Common Issues**
   - Transfer ID must be unique
   - From account must have sufficient balance
   - Processing time is simulated (1-5 seconds)
   - Success rate is 80% (20% chance of failure)

## 🧹 Clean Up

To stop and clean up all services:
```bash
docker compose down -v
```

This will remove all containers and volumes, including the database data.

To stop and clean up all services:
```bash
docker-compose down
```
