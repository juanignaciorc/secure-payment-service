package mocks

import (
	"context"
	"time"

	"secure-payment-service/internal/models"
	"secure-payment-service/internal/repositories/interfaces"
	"github.com/stretchr/testify/mock"
)

type TransferRepositoryMock struct {
	mock.Mock
}

// Ensure TransferRepositoryMock implements interfaces.TransferRepository
var _ interfaces.TransferRepository = (*TransferRepositoryMock)(nil)

func (m *TransferRepositoryMock) CreateTransfer(ctx context.Context, transfer *models.Transfer) error {
	args := m.Called(ctx, transfer)
	return args.Error(0)
}

func (m *TransferRepositoryMock) GetPendingTransfers(ctx context.Context, timeout time.Duration) ([]*models.Transfer, error) {
	args := m.Called(ctx, timeout)
	return args.Get(0).([]*models.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) GetTransfer(ctx context.Context, id string) (*models.Transfer, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Transfer), args.Error(1)
}

func (m *TransferRepositoryMock) UpdateTransferStatus(ctx context.Context, id string, status models.TransferStatus) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}
