package usecases

import (
	"context"
	"time"

	"manager/internal/repositories"
)

type DashboardSummary struct {
	TotalBalance     int64           `json:"total_balance"`
	MonthlyIncome    int64           `json:"monthly_income"`
	MonthlyExpense   int64           `json:"monthly_expense"`
	CategoryExpenses map[int64]int64 `json:"category_expenses"`
}

type DashboardUsecase struct {
	transactionRepo *repositories.TransactionRepository
	categoryRepo    *repositories.CategoryRepository
}

func NewDashboardUsecase(transactionRepo *repositories.TransactionRepository, categoryRepo *repositories.CategoryRepository) *DashboardUsecase {
	return &DashboardUsecase{
		transactionRepo: transactionRepo,
		categoryRepo:    categoryRepo,
	}
}

func (u *DashboardUsecase) GetSummary(ctx context.Context) (*DashboardSummary, error) {
	now := time.Now()

	// Buscar saldo total
	totalBalance, err := u.transactionRepo.GetTotalBalance(ctx)
	if err != nil {
		return nil, err
	}

	// Buscar receitas e despesas do mês
	monthlyIncome, monthlyExpense, err := u.transactionRepo.GetMonthlySummary(ctx, now.Year(), int(now.Month()))
	if err != nil {
		return nil, err
	}

	// Buscar gastos por categoria
	categoryExpenses, err := u.transactionRepo.GetCategoryExpenses(ctx, now.Year(), int(now.Month()))
	if err != nil {
		return nil, err
	}

	return &DashboardSummary{
		TotalBalance:     totalBalance,
		MonthlyIncome:    monthlyIncome,
		MonthlyExpense:   monthlyExpense,
		CategoryExpenses: categoryExpenses,
	}, nil
}
