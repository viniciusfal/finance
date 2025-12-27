package entity

import "time"

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusPaid      TransactionStatus = "paid"
	TransactionStatusOverdue   TransactionStatus = "overdue"
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

type Transaction struct {
	ID                int64             `json:"id"`
	Title             string            `json:"title"`
	Description       *string           `json:"description,omitempty"`
	AmountCents       int64             `json:"amount_cents"`
	Type              TransactionType   `json:"type"`
	CategoryID        *int64            `json:"category_id,omitempty"`
	DueDate           time.Time         `json:"due_date"`
	IsRecurring       bool              `json:"is_recurring"`
	IsInstallment     bool              `json:"is_installment"`
	TotalInstallments int               `json:"total_installments"`
	Status            TransactionStatus `json:"status"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
	Category          *Category         `json:"category,omitempty"`
	Installments      []Installment     `json:"installments,omitempty"`
}
