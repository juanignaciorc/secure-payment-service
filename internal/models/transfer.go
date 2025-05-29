package models

import "time"

type TransferStatus string

const (
	TransferStatusPending   TransferStatus = "PENDING"
	TransferStatusCompleted TransferStatus = "COMPLETED"
	TransferStatusFailed    TransferStatus = "FAILED"
)

type Transfer struct {
	ID          string        `json:"id"`
	FromAccount string        `json:"from_account"`
	ToAccount   string        `json:"to_account"`
	Amount      float64       `json:"amount"`
	Status      TransferStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
