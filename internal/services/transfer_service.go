package services

import (
	"context"
	"errors"
	"time"

	"secure-payment-service/internal/models"
	"secure-payment-service/internal/repositories/interfaces"
)

type TransferService struct {
	accountRepo  interfaces.AccountRepository
	transferRepo interfaces.TransferRepository
}

func (s *TransferService) AccountRepo() interfaces.AccountRepository {
	return s.accountRepo
}

func NewTransferService(accountRepo interfaces.AccountRepository, transferRepo interfaces.TransferRepository) *TransferService {
	return &TransferService{
		accountRepo:  accountRepo,
		transferRepo: transferRepo,
	}
}

func (s *TransferService) CreateTransfer(ctx context.Context, transfer *models.Transfer) error {
	// Validar que las cuentas existan
	fromAccount, err := s.accountRepo.GetAccount(ctx, transfer.FromAccount)
	if err != nil {
		return err
	}

	// Verificar que la cuenta destino exista
	if _, err := s.accountRepo.GetAccount(ctx, transfer.ToAccount); err != nil {
		return err
	}

	// Verificar saldo suficiente
	if fromAccount.Balance < transfer.Amount {
		return errors.New("insufficient funds")
	}

	// Crear transferencia con estado PENDING
	transfer.ID = time.Now().Format("20060102150405")
	transfer.Status = models.TransferStatusPending
	transfer.CreatedAt = time.Now()
	transfer.UpdatedAt = time.Now()

	return s.transferRepo.CreateTransfer(ctx, transfer)
}

func (s *TransferService) GetPendingTransfers(ctx context.Context, timeout time.Duration) ([]*models.Transfer, error) {
	return s.transferRepo.GetPendingTransfers(ctx, timeout)
}

func (s *TransferService) UpdateTransferStatus(ctx context.Context, id string, status models.TransferStatus) error {
	transfer, err := s.transferRepo.GetTransfer(ctx, id)
	if err != nil {
		return err
	}

	// Solo se pueden actualizar transferencias en estado PENDING
	if transfer.Status != models.TransferStatusPending {
		return errors.New("transfer is not in pending state")
	}

	// Si es COMPLETED, actualizar balances
	if status == models.TransferStatusCompleted {
		// Decrementar saldo de la cuenta origen
		if err := s.accountRepo.UpdateBalance(ctx, transfer.FromAccount, transfer.Amount*-1); err != nil {
			return err
		}

		// Incrementar saldo de la cuenta destino
		if err := s.accountRepo.UpdateBalance(ctx, transfer.ToAccount, transfer.Amount); err != nil {
			return err
		}
	}

	transfer.Status = status
	transfer.UpdatedAt = time.Now()

	return s.transferRepo.UpdateTransferStatus(ctx, transfer.ID, transfer.Status)
}

func (s *TransferService) GetTransfer(ctx context.Context, id string) (*models.Transfer, error) {
	return s.transferRepo.GetTransfer(ctx, id)
}

func (s *TransferService) GetAccountBalance(ctx context.Context, id string) (float64, error) {
	account, err := s.accountRepo.GetAccount(ctx, id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
