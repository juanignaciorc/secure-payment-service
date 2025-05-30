FROM golang:1.22-alpine AS base

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build server image
FROM base AS server
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go && \
    chmod +x server
CMD ["./server"]

# Build simulator image
FROM base AS simulator
RUN CGO_ENABLED=0 GOOS=linux go build -o simulator ./cmd/simulator/main.go && \
    chmod +x simulator
CMD ["./simulator"]
