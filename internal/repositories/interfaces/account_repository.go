package interfaces

import (
	"context"
	"secure-payment-service/internal/models"
)

type AccountRepository interface {
	GetAccount(ctx context.Context, id string) (*models.Account, error)
	UpdateBalance(ctx context.Context, id string, amount float64) error
	CreateAccount(ctx context.Context, account *models.Account) error
}
