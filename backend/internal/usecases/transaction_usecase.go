package usecases

import (
	"context"
	"fmt"

	"manager/internal/entity"
	"manager/internal/repositories"
)

type TransactionUsecase struct {
	repo *repositories.TransactionRepository
}

func NewTransactionUsecase(repo *repositories.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{repo: repo}
}

func (u *TransactionUsecase) Create(ctx context.Context, transaction *entity.Transaction) error {
	if transaction.Title == "" {
		return fmt.Errorf("transaction title is required")
	}

	if transaction.AmountCents <= 0 {
		return fmt.Errorf("transaction amount must be greater than zero")
	}

	if transaction.Type != entity.TransactionTypeIncome && transaction.Type != entity.TransactionTypeExpense {
		return fmt.Errorf("invalid transaction type")
	}

	if transaction.Status == "" {
		transaction.Status = entity.TransactionStatusPending
	}

	if transaction.IsInstallment {
		if transaction.TotalInstallments <= 1 {
			return fmt.Errorf("installment transactions must have more than 1 installment")
		}
	} else {
		transaction.TotalInstallments = 1
	}

	// Criar transação
	if err := u.repo.Create(ctx, transaction); err != nil {
		return err
	}

	// Se for parcelada, criar as parcelas
	if transaction.IsInstallment {
		if err := u.createInstallments(ctx, transaction); err != nil {
			return err
		}
	}

	return nil
}

func (u *TransactionUsecase) createInstallments(ctx context.Context, transaction *entity.Transaction) error {
	// Calcular valor por parcela
	amountPerInstallment := transaction.AmountCents / int64(transaction.TotalInstallments)
	remainder := transaction.AmountCents % int64(transaction.TotalInstallments)

	// Criar cada parcela com intervalo de 30 dias
	for i := 1; i <= transaction.TotalInstallments; i++ {
		installmentAmount := amountPerInstallment
		// Adicionar resto na primeira parcela
		if i == 1 {
			installmentAmount += remainder
		}

		// Calcular data: data inicial + (30 dias × (número da parcela - 1))
		dueDate := transaction.DueDate.AddDate(0, 0, 30*(i-1))

		installment := &entity.Installment{
			TransactionID:     transaction.ID,
			InstallmentNumber: i,
			AmountCents:       installmentAmount,
			DueDate:           dueDate,
			Status:            entity.InstallmentStatusPending,
		}

		if err := u.repo.CreateInstallment(ctx, installment); err != nil {
			return fmt.Errorf("failed to create installment %d: %w", i, err)
		}
	}

	return nil
}

func (u *TransactionUsecase) GetAll(ctx context.Context) ([]entity.Transaction, error) {
	return u.repo.GetAll(ctx)
}

func (u *TransactionUsecase) GetByID(ctx context.Context, id int64) (*entity.Transaction, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid transaction id")
	}

	return u.repo.GetByID(ctx, id)
}

func (u *TransactionUsecase) Update(ctx context.Context, transaction *entity.Transaction) error {
	if transaction.ID <= 0 {
		return fmt.Errorf("invalid transaction id")
	}

	if transaction.Title == "" {
		return fmt.Errorf("transaction title is required")
	}

	if transaction.AmountCents <= 0 {
		return fmt.Errorf("transaction amount must be greater than zero")
	}

	return u.repo.Update(ctx, transaction)
}

func (u *TransactionUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid transaction id")
	}

	return u.repo.Delete(ctx, id)
}

func (u *TransactionUsecase) PayInstallment(ctx context.Context, transactionID int64, installmentNumber int) error {
	if transactionID <= 0 {
		return fmt.Errorf("invalid transaction id")
	}

	if installmentNumber <= 0 {
		return fmt.Errorf("invalid installment number")
	}

	return u.repo.PayInstallment(ctx, transactionID, installmentNumber)
}

