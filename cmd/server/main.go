package main

import (
	"context"
	"net/http"

	"secure-payment-service/internal/config"
	"secure-payment-service/internal/handlers"
	"secure-payment-service/internal/logging"
	"secure-payment-service/internal/migrations"
	"secure-payment-service/internal/repositories"
	"secure-payment-service/internal/services"
	"secure-payment-service/internal/workers"
	"time"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger := logging.NewLogger(logging.LevelError)
		logger.Fatal("Failed to load configuration:", err)
	}

	// Initialize logger with configured level
	logger := logging.NewLogger(logging.LevelInfo)
	logger.Info("Configuration loaded successfully")

	// Apply migrations
	if err := migrations.ApplyMigrations(cfg.DB); err != nil {
		logger.Fatal("Failed to apply migrations:", err)
	}

	// Initialize repositories
	accountRepo := repositories.NewAccountRepository(cfg.DB)
	transferRepo := repositories.NewTransferRepository(cfg.DB)

	// Initialize services
	transferService := services.NewTransferService(accountRepo, transferRepo)

	// Initialize handlers
	transferHandler := handlers.NewTransferHandler(transferService, accountRepo)
	webhookHandler := handlers.NewWebhookHandler(transferService, logger)

	// Initialize transfer status checker
	statusChecker := workers.NewTransferStatusChecker(
		transferService,
		time.Minute, // Check every minute
		time.Hour,   // Timeout after 1 hour
	)

	// Create router
	mux := http.NewServeMux()

	// Transfer endpoints
	mux.HandleFunc("/api/transfers", transferHandler.CreateTransfer)
	mux.HandleFunc("/api/transfers/status", transferHandler.UpdateTransferStatus)
	mux.HandleFunc("/api/transfers/info", transferHandler.GetTransfer)
	mux.HandleFunc("/api/accounts", transferHandler.CreateAccount)
	mux.HandleFunc("/api/accounts/balance", transferHandler.GetAccountBalance)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/api/webhook/transfer", webhookHandler.HandleWebhook)

	// Start server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	// Start status checker in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go statusChecker.Start(ctx)

	logger.Info("Server starting on port", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.WithError(err).Fatal("Server failed")
	}
}
