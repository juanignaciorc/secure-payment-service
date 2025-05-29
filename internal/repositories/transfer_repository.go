package repositories

import (
	"context"
	"database/sql"
	"time"

	"secure-payment-service/internal/models"
)

type TransferRepository struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{db: db}
}

func (r *TransferRepository) CreateTransfer(ctx context.Context, transfer *models.Transfer) error {
	query := `
		INSERT INTO transfers (id, from_account, to_account, amount, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		transfer.ID,
		transfer.FromAccount,
		transfer.ToAccount,
		transfer.Amount,
		transfer.Status,
		transfer.CreatedAt,
		transfer.UpdatedAt,
	)
	return err
}

func (r *TransferRepository) GetPendingTransfers(ctx context.Context, timeout time.Duration) ([]*models.Transfer, error) {
	query := `
		SELECT id, from_account, to_account, amount, status, created_at, updated_at
		FROM transfers
		WHERE status = $1
		AND created_at < $2
	`
	var transfers []*models.Transfer
	rows, err := r.db.QueryContext(ctx, query,
		models.TransferStatusPending,
		time.Now().Add(-timeout),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transfer models.Transfer
		err := rows.Scan(
			&transfer.ID,
			&transfer.FromAccount,
			&transfer.ToAccount,
			&transfer.Amount,
			&transfer.Status,
			&transfer.CreatedAt,
			&transfer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, &transfer)
	}
	return transfers, nil
}

func (r *TransferRepository) GetTransfer(ctx context.Context, id string) (*models.Transfer, error) {
	query := `
		SELECT id, from_account, to_account, amount, status, created_at, updated_at
		FROM transfers
		WHERE id = $1
	`
	var transfer models.Transfer
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(
		&transfer.ID,
		&transfer.FromAccount,
		&transfer.ToAccount,
		&transfer.Amount,
		&transfer.Status,
		&transfer.CreatedAt,
		&transfer.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &transfer, nil
}

func (r *TransferRepository) UpdateTransferStatus(ctx context.Context, id string, status models.TransferStatus) error {
	query := `
		UPDATE transfers
		SET status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.db.ExecContext(ctx, query, status, time.Now(), id)
	return err
}
