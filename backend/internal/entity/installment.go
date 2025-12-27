package entity

import "time"

type InstallmentStatus string

const (
	InstallmentStatusPending   InstallmentStatus = "pending"
	InstallmentStatusPaid      InstallmentStatus = "paid"
	InstallmentStatusOverdue   InstallmentStatus = "overdue"
	InstallmentStatusCancelled InstallmentStatus = "cancelled"
)

type Installment struct {
	ID                int64             `json:"id"`
	TransactionID     int64             `json:"transaction_id"`
	InstallmentNumber int               `json:"installment_number"`
	AmountCents       int64             `json:"amount_cents"`
	DueDate           time.Time         `json:"due_date"`
	Status            InstallmentStatus `json:"status"`
	PaidAt            *time.Time        `json:"paid_at,omitempty"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}
