package interfaces

import (
	"context"
	"time"

	"secure-payment-service/internal/models"
)

type TransferRepository interface {
	CreateTransfer(ctx context.Context, transfer *models.Transfer) error
	GetPendingTransfers(ctx context.Context, timeout time.Duration) ([]*models.Transfer, error)
	GetTransfer(ctx context.Context, id string) (*models.Transfer, error)
	UpdateTransferStatus(ctx context.Context, id string, status models.TransferStatus) error
}
