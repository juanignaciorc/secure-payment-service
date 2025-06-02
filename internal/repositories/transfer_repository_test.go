package repositories

import (
	"context"
	"testing"
	"time"

	"secure-payment-service/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTransferRepository_CreateTransfer(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewTransferRepository(db)

	transfer := &models.Transfer{
		ID:          "1",
		FromAccount: "123",
		ToAccount:   "456",
		Amount:      100.0,
		Status:      models.TransferStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec("INSERT INTO transfers").
		WithArgs(
			transfer.ID,
			transfer.FromAccount,
			transfer.ToAccount,
			transfer.Amount,
			transfer.Status,
			transfer.CreatedAt,
			transfer.UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateTransfer(ctx, transfer)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestTransferRepository_GetPendingTransfers(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewTransferRepository(db)

	timeout := time.Hour
	transfers := []*models.Transfer{
		{
			ID:          "1",
			FromAccount: "123",
			ToAccount:   "456",
			Amount:      100.0,
			Status:      models.TransferStatusPending,
			CreatedAt:   time.Now().Add(-2 * time.Hour),
			UpdatedAt:   time.Now().Add(-2 * time.Hour),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "from_account", "to_account", "amount", "status", "created_at", "updated_at"}).
		AddRow(
			transfers[0].ID,
			transfers[0].FromAccount,
			transfers[0].ToAccount,
			transfers[0].Amount,
			transfers[0].Status,
			transfers[0].CreatedAt,
			transfers[0].UpdatedAt,
		)

	mock.ExpectQuery("SELECT id, from_account, to_account, amount, status, created_at, updated_at").
		WithArgs(models.TransferStatusPending, time.Now().Add(-timeout)).
		WillReturnRows(rows)

	got, err := repo.GetPendingTransfers(ctx, timeout)
	assert.NoError(t, err)
	assert.Equal(t, transfers, got)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestTransferRepository_GetTransfer(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewTransferRepository(db)

	id := "1"
	transfer := &models.Transfer{
		ID:          id,
		FromAccount: "123",
		ToAccount:   "456",
		Amount:      100.0,
		Status:      models.TransferStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "from_account", "to_account", "amount", "status", "created_at", "updated_at"}).
		AddRow(
			transfer.ID,
			transfer.FromAccount,
			transfer.ToAccount,
			transfer.Amount,
			transfer.Status,
			transfer.CreatedAt,
			transfer.UpdatedAt,
		)

	mock.ExpectQuery("SELECT id, from_account, to_account, amount, status, created_at, updated_at").
		WithArgs(id).
		WillReturnRows(rows)

	got, err := repo.GetTransfer(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, transfer, got)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestTransferRepository_UpdateTransferStatus(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewTransferRepository(db)

	id := "1"
	status := models.TransferStatusCompleted

	mock.ExpectExec("UPDATE transfers").
		WithArgs(status, time.Now(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateTransferStatus(ctx, id, status)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
