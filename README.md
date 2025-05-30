# Secure Payment Service

A microservice for handling secure bank account transfers with Go.

## 🚀 Quick Start

1. **Prerequisites**
   - Docker installed

2. **Build and Run Services**
   ```bash
   # Build the service
   docker build -t payment-service .
   
   # Run the database container
   docker run --name payment-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=payment_db -p 5432:5432 -d postgres
   
   # Run the service
   docker run --name payment-service --link payment-db:db -p 8080:8080 -d payment-service
   ```

3. **Verify Services**
   - Payment service: `http://localhost:8080`
   - Database: `postgres://postgres:postgres@localhost:5432/payment_db`

## 🧪 Complete Test Flow

### 1. Initial Setup
- All services will start automatically with Docker
- Database is initialized with required tables
- No manual setup needed

### 2. Create Test Accounts
```bash
# Create account 1 (sender)
curl -X POST http://localhost:8080/api/accounts \
  -H "Content-Type: application/json" \
  -d '{"id": "acc-100", "balance": 1000.00}'

# Create account 2 (receiver)
curl -X POST http://localhost:8080/api/accounts \
  -H "Content-Type: application/json" \
  -d '{"id": "acc-200", "balance": 500.00}'
```

### 3. Perform a Transfer
```bash
# Create a transfer request
curl -X POST http://localhost:8080/api/transfers \
  -H "Content-Type: application/json" \
  -d '{"id": "transfer-001", "from_account": "acc-100", "to_account": "acc-200", "amount": 200.00}'
```

### 4. Verify Transfer Status
```bash
# Check transfer status
curl "http://localhost:8080/api/transfers/info?id=transfer-001"

# Expected response: {"id":"transfer-001","status":"PENDING"}
```

### 5. Update Transfer Status via Simulator
```bash
# Update transfer status to COMPLETED
curl "http://localhost:8081/api/transfers/status?id=transfer-001&status=COMPLETED"
```

### 6. Verify Final Balances
```bash
# Check sender's balance
curl "http://localhost:8080/api/accounts/balance?account_id=acc-100"
# Expected: 800.00

# Check receiver's balance
curl "http://localhost:8080/api/accounts/balance?account_id=acc-200"
# Expected: 700.00
```

## 🛠️ Troubleshooting

1. **If services don't start**
   - Ensure Docker and Docker Compose are installed
   - Run `docker-compose down` to clean up previous containers
   - Try rebuilding with `docker-compose up --build`

2. **If database connection fails**
   - Wait a few seconds for database initialization
   - Check if PostgreSQL is running with `docker ps`

3. **If transfer fails**
   - Verify account balances are sufficient
   - Check transfer ID format
   - Ensure accounts exist

## 📝 API Documentation

### Accounts
- `POST /api/accounts`: Create new account
- `GET /api/accounts/balance`: Get account balance

### Transfers
- `POST /api/transfers`: Create new transfer
- `GET /api/transfers/info`: Get transfer status
- `GET /api/transfers/status`: Update transfer status (via simulator)

## 🧹 Clean Up

To stop and clean up all services:
```bash
docker-compose down
```
