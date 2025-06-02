package mocks

import (
	"context"

	"secure-payment-service/internal/models"
	"secure-payment-service/internal/repositories/interfaces"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) GetAccount(ctx context.Context, id string) (*models.Account, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *AccountRepositoryMock) UpdateBalance(ctx context.Context, id string, amount float64) error {
	args := m.Called(ctx, id, amount)
	return args.Error(0)
}

func (m *AccountRepositoryMock) CreateAccount(ctx context.Context, account *models.Account) error {
	args := m.Called(ctx, account)
	return args.Error(0)
}

// Make sure the mock implements the repository interface
var _ interfaces.AccountRepository = (*AccountRepositoryMock)(nil)
