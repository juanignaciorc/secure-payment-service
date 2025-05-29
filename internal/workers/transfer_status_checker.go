package workers

import (
	"context"
	"time"

	"secure-payment-service/internal/models"
	"secure-payment-service/internal/services"
)

type TransferStatusChecker struct {
	transferService *services.TransferService
	checkInterval   time.Duration
	timeout         time.Duration
}

func NewTransferStatusChecker(transferService *services.TransferService, checkInterval time.Duration, timeout time.Duration) *TransferStatusChecker {
	return &TransferStatusChecker{
		transferService: transferService,
		checkInterval:   checkInterval,
		timeout:         timeout,
	}
}

func (c *TransferStatusChecker) Start(ctx context.Context) {
	ticker := time.NewTicker(c.checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.checkPendingTransfers(ctx)
		}
	}
}

func (c *TransferStatusChecker) checkPendingTransfers(ctx context.Context) {
	// Obtener transferencias pendientes que superen el timeout
	transfers, err := c.transferService.GetPendingTransfers(ctx, c.timeout)
	if err != nil {
		// Log the error but continue
		return
	}

	for _, transfer := range transfers {
		// Aquí podrías implementar la lógica de reintentar el pago o marcar como failed
		// Por ejemplo:
		if err := c.transferService.UpdateTransferStatus(ctx, transfer.ID, models.TransferStatusFailed); err != nil {
			// Log the error
		}
	}
}
