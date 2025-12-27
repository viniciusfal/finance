package repositories

import (
	"context"
	"fmt"
	"time"

	"manager/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *entity.Transaction) error {
	query := `
		INSERT INTO transactions (title, description, amount_cents, type, category_id, due_date, is_recurring, is_installment, total_installments, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		transaction.Title,
		transaction.Description,
		transaction.AmountCents,
		transaction.Type,
		transaction.CategoryID,
		transaction.DueDate,
		transaction.IsRecurring,
		transaction.IsInstallment,
		transaction.TotalInstallments,
		transaction.Status,
	).Scan(&transaction.ID, &transaction.CreatedAt, &transaction.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return nil
}

func (r *TransactionRepository) GetAll(ctx context.Context) ([]entity.Transaction, error) {
	query := `
		SELECT 
			t.id, t.title, t.description, t.amount_cents, t.type, 
			t.category_id, t.due_date, t.is_recurring, t.is_installment, 
			t.total_installments, t.status, t.created_at, t.updated_at,
			c.id, c.name, c.description, c.color, c.icon, c.created_at, c.updated_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		ORDER BY t.due_date DESC, t.created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		var t entity.Transaction
		var categoryID *int64
		var cat entity.Category
		var catID *int64

		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.AmountCents,
			&t.Type,
			&categoryID,
			&t.DueDate,
			&t.IsRecurring,
			&t.IsInstallment,
			&t.TotalInstallments,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
			&catID,
			&cat.Name,
			&cat.Description,
			&cat.Color,
			&cat.Icon,
			&cat.CreatedAt,
			&cat.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}

		t.CategoryID = categoryID
		if catID != nil {
			cat.ID = *catID
			t.Category = &cat
		}

		// Buscar parcelas se for parcelada
		if t.IsInstallment {
			installments, err := r.GetInstallmentsByTransactionID(ctx, t.ID)
			if err == nil {
				t.Installments = installments
			}
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepository) GetByID(ctx context.Context, id int64) (*entity.Transaction, error) {
	query := `
		SELECT 
			t.id, t.title, t.description, t.amount_cents, t.type, 
			t.category_id, t.due_date, t.is_recurring, t.is_installment, 
			t.total_installments, t.status, t.created_at, t.updated_at,
			c.id, c.name, c.description, c.color, c.icon, c.created_at, c.updated_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.id = $1
	`

	var t entity.Transaction
	var categoryID *int64
	var cat entity.Category
	var catID *int64

	err := r.db.QueryRow(ctx, query, id).Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.AmountCents,
		&t.Type,
		&categoryID,
		&t.DueDate,
		&t.IsRecurring,
		&t.IsInstallment,
		&t.TotalInstallments,
		&t.Status,
		&t.CreatedAt,
		&t.UpdatedAt,
		&catID,
		&cat.Name,
		&cat.Description,
		&cat.Color,
		&cat.Icon,
		&cat.CreatedAt,
		&cat.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	t.CategoryID = categoryID
	if catID != nil {
		cat.ID = *catID
		t.Category = &cat
	}

	// Buscar parcelas se for parcelada
	if t.IsInstallment {
		installments, err := r.GetInstallmentsByTransactionID(ctx, t.ID)
		if err == nil {
			t.Installments = installments
		}
	}

	return &t, nil
}

func (r *TransactionRepository) Update(ctx context.Context, transaction *entity.Transaction) error {
	query := `
		UPDATE transactions
		SET title = $1, description = $2, amount_cents = $3, type = $4, 
		    category_id = $5, due_date = $6, is_recurring = $7, 
		    is_installment = $8, total_installments = $9, status = $10, updated_at = NOW()
		WHERE id = $11
		RETURNING updated_at
	`

	err := r.db.QueryRow(ctx, query,
		transaction.Title,
		transaction.Description,
		transaction.AmountCents,
		transaction.Type,
		transaction.CategoryID,
		transaction.DueDate,
		transaction.IsRecurring,
		transaction.IsInstallment,
		transaction.TotalInstallments,
		transaction.Status,
		transaction.ID,
	).Scan(&transaction.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	return nil
}

func (r *TransactionRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM transactions WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete transaction: %w", err)
	}

	return nil
}

func (r *TransactionRepository) CreateInstallment(ctx context.Context, installment *entity.Installment) error {
	query := `
		INSERT INTO transaction_installments (transaction_id, installment_number, amount_cents, due_date, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		installment.TransactionID,
		installment.InstallmentNumber,
		installment.AmountCents,
		installment.DueDate,
		installment.Status,
	).Scan(&installment.ID, &installment.CreatedAt, &installment.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create installment: %w", err)
	}

	return nil
}

func (r *TransactionRepository) GetInstallmentsByTransactionID(ctx context.Context, transactionID int64) ([]entity.Installment, error) {
	query := `
		SELECT id, transaction_id, installment_number, amount_cents, due_date, status, paid_at, created_at, updated_at
		FROM transaction_installments
		WHERE transaction_id = $1
		ORDER BY installment_number ASC
	`

	rows, err := r.db.Query(ctx, query, transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get installments: %w", err)
	}
	defer rows.Close()

	var installments []entity.Installment
	for rows.Next() {
		var inst entity.Installment
		err := rows.Scan(
			&inst.ID,
			&inst.TransactionID,
			&inst.InstallmentNumber,
			&inst.AmountCents,
			&inst.DueDate,
			&inst.Status,
			&inst.PaidAt,
			&inst.CreatedAt,
			&inst.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}
		installments = append(installments, inst)
	}

	return installments, nil
}

func (r *TransactionRepository) PayInstallment(ctx context.Context, transactionID int64, installmentNumber int) error {
	query := `
		UPDATE transaction_installments
		SET status = $1, paid_at = NOW(), updated_at = NOW()
		WHERE transaction_id = $2 AND installment_number = $3
	`

	_, err := r.db.Exec(ctx, query, entity.InstallmentStatusPaid, transactionID, installmentNumber)
	if err != nil {
		return fmt.Errorf("failed to pay installment: %w", err)
	}

	return nil
}

func (r *TransactionRepository) GetMonthlySummary(ctx context.Context, year int, month int) (int64, int64, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	// Para transações não parceladas, usar o valor total
	// Para transações parceladas, somar apenas as parcelas com vencimento no mês
	query := `
		SELECT 
			COALESCE(SUM(
				CASE 
					WHEN t.is_installment = false AND t.type = 'income' THEN t.amount_cents
					WHEN t.is_installment = true AND t.type = 'income' THEN COALESCE(SUM(ti.amount_cents), 0)
					ELSE 0
				END
			), 0) as income,
			COALESCE(SUM(
				CASE 
					WHEN t.is_installment = false AND t.type = 'expense' THEN t.amount_cents
					WHEN t.is_installment = true AND t.type = 'expense' THEN COALESCE(SUM(ti.amount_cents), 0)
					ELSE 0
				END
			), 0) as expense
		FROM transactions t
		LEFT JOIN transaction_installments ti ON t.id = ti.transaction_id 
			AND ti.due_date >= $1 AND ti.due_date < $2
		WHERE (t.is_installment = false AND t.due_date >= $1 AND t.due_date < $2)
		   OR (t.is_installment = true)
		AND t.status != 'cancelled'
		GROUP BY t.id
	`

	// Query simplificada - somar transações não parceladas e parcelas do mês
	query = `
		SELECT 
			COALESCE(SUM(CASE WHEN type = 'income' AND is_installment = false THEN amount_cents ELSE 0 END), 0) +
			COALESCE(SUM(CASE WHEN type = 'income' AND is_installment = true THEN 0 ELSE 0 END), 0) +
			COALESCE((SELECT SUM(amount_cents) FROM transaction_installments ti 
				INNER JOIN transactions t2 ON ti.transaction_id = t2.id 
				WHERE ti.due_date >= $1 AND ti.due_date < $2 AND t2.type = 'income'), 0) as income,
			COALESCE(SUM(CASE WHEN type = 'expense' AND is_installment = false THEN amount_cents ELSE 0 END), 0) +
			COALESCE((SELECT SUM(amount_cents) FROM transaction_installments ti 
				INNER JOIN transactions t2 ON ti.transaction_id = t2.id 
				WHERE ti.due_date >= $1 AND ti.due_date < $2 AND t2.type = 'expense'), 0) as expense
		FROM transactions
		WHERE (is_installment = false AND due_date >= $1 AND due_date < $2)
		AND status != 'cancelled'
	`

	var income, expense int64
	err := r.db.QueryRow(ctx, query, startDate, endDate).Scan(&income, &expense)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get monthly summary: %w", err)
	}

	return income, expense, nil
}

func (r *TransactionRepository) GetTotalBalance(ctx context.Context) (int64, error) {
	// Para transações não parceladas, usar o valor total
	// Para transações parceladas, somar apenas as parcelas pagas
	query := `
		SELECT 
			COALESCE(SUM(
				CASE 
					WHEN t.is_installment = false THEN 
						CASE WHEN t.type = 'income' THEN t.amount_cents ELSE -t.amount_cents END
					ELSE 0
				END
			), 0) +
			COALESCE(SUM(
				CASE 
					WHEN t.is_installment = true THEN 
						CASE WHEN t.type = 'income' THEN ti.amount_cents ELSE -ti.amount_cents END
					ELSE 0
				END
			), 0) as balance
		FROM transactions t
		LEFT JOIN transaction_installments ti ON t.id = ti.transaction_id AND ti.status = 'paid'
		WHERE t.status != 'cancelled'
	`

	var balance int64
	err := r.db.QueryRow(ctx, query).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("failed to get total balance: %w", err)
	}

	return balance, nil
}

func (r *TransactionRepository) GetCategoryExpenses(ctx context.Context, year int, month int) (map[int64]int64, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	query := `
		SELECT category_id, COALESCE(SUM(amount_cents), 0) as total
		FROM transactions
		WHERE type = 'expense' 
		AND due_date >= $1 AND due_date < $2
		AND status != 'cancelled'
		AND category_id IS NOT NULL
		GROUP BY category_id
	`

	rows, err := r.db.Query(ctx, query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get category expenses: %w", err)
	}
	defer rows.Close()

	expenses := make(map[int64]int64)
	for rows.Next() {
		var categoryID int64
		var total int64
		if err := rows.Scan(&categoryID, &total); err != nil {
			return nil, fmt.Errorf("failed to scan category expense: %w", err)
		}
		expenses[categoryID] = total
	}

	return expenses, nil
}

