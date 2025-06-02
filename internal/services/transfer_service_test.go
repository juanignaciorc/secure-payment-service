package services

import (
	"context"
	"testing"
	"time"

	"secure-payment-service/internal/models"
	"secure-payment-service/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransferService_CreateTransfer(t *testing.T) {
	ctx := context.Background()
	mockAccountRepo := &mocks.AccountRepositoryMock{}
	mockTransferRepo := &mocks.TransferRepositoryMock{}

	// Set up mock expectations
	mockAccountRepo.On("GetAccount", ctx, "456").Return(&models.Account{
		ID:      "456",
		Balance: 1000.0,
	}, nil)
	mockAccountRepo.On("GetAccount", ctx, "789").Return(&models.Account{
		ID:      "789",
		Balance: 500.0,
	}, nil)
	mockTransferRepo.On("CreateTransfer", ctx, mock.AnythingOfType("*models.Transfer")).Return(nil)

	service := NewTransferService(mockAccountRepo, mockTransferRepo)

	tests := []struct {
		name    string
		transfer *models.Transfer
		wantErr  bool
	}{
		{
			name: "valid transfer",
			transfer: &models.Transfer{
				ID:           "123",
				FromAccount:   "456",
				ToAccount:     "789",
				Amount:       100.0,
				Status:       models.TransferStatusPending,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CreateTransfer(ctx, tt.transfer)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransfer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.NotNil(t, tt.transfer.ID)
				assert.Equal(t, models.TransferStatusPending, tt.transfer.Status)
				assert.NotZero(t, tt.transfer.CreatedAt)
				assert.NotZero(t, tt.transfer.UpdatedAt)
			}
		})
	}
}

func TestTransferService_UpdateTransferStatus(t *testing.T) {
	ctx := context.Background()
	mockAccountRepo := &mocks.AccountRepositoryMock{}
	mockTransferRepo := &mocks.TransferRepositoryMock{}

	// Set up mock expectations
	mockTransferRepo.On("GetTransfer", ctx, "123").Return(&models.Transfer{
		ID:           "123",
		FromAccount:  "456",
		ToAccount:    "789",
		Amount:       100.0,
		Status:       models.TransferStatusPending,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil)
	mockTransferRepo.On("UpdateTransferStatus", ctx, "123", models.TransferStatusCompleted).Return(nil)

	// Mock account balance updates
	mockAccountRepo.On("UpdateBalance", ctx, "456", -100.0).Return(nil)
	mockAccountRepo.On("UpdateBalance", ctx, "789", 100.0).Return(nil)

	service := NewTransferService(mockAccountRepo, mockTransferRepo)

	tests := []struct {
		name    string
		id      string
		status  models.TransferStatus
		wantErr bool
	}{
		{
			name:    "valid status update",
			id:      "123",
			status:  models.TransferStatusCompleted,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := service.UpdateTransferStatus(ctx, tt.id, tt.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTransferStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
