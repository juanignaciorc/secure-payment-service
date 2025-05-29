package repositories

import (
	"context"
	"database/sql"
	"secure-payment-service/internal/models"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *models.Account) error {
	query := `
		INSERT INTO accounts (id, balance, created_at)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, query, account.ID, account.Balance, account.CreatedAt)
	return err
}

func (r *AccountRepository) GetAccount(ctx context.Context, id string) (*models.Account, error) {
	query := `
		SELECT id, balance, created_at
		FROM accounts
		WHERE id = $1
	`
	var account models.Account
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&account.ID, &account.Balance, &account.CreatedAt); err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *AccountRepository) UpdateBalance(ctx context.Context, id string, amount float64) error {
	query := `
		UPDATE accounts
		SET balance = balance + $2
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id, amount)
	return err
}
