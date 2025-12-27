package usecases

import (
	"context"
	"fmt"

	"manager/internal/entity"
	"manager/internal/repositories"
)

type CategoryUsecase struct {
	repo *repositories.CategoryRepository
}

func NewCategoryUsecase(repo *repositories.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) Create(ctx context.Context, category *entity.Category) error {
	if category.Name == "" {
		return fmt.Errorf("category name is required")
	}

	if category.Color == "" {
		category.Color = "#6366F1"
	}

	if category.Icon == "" {
		category.Icon = "tag"
	}

	return u.repo.Create(ctx, category)
}

func (u *CategoryUsecase) GetAll(ctx context.Context) ([]entity.Category, error) {
	return u.repo.GetAll(ctx)
}

func (u *CategoryUsecase) GetByID(ctx context.Context, id int64) (*entity.Category, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid category id")
	}

	return u.repo.GetByID(ctx, id)
}

func (u *CategoryUsecase) Update(ctx context.Context, category *entity.Category) error {
	if category.ID <= 0 {
		return fmt.Errorf("invalid category id")
	}

	if category.Name == "" {
		return fmt.Errorf("category name is required")
	}

	return u.repo.Update(ctx, category)
}

func (u *CategoryUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid category id")
	}

	return u.repo.Delete(ctx, id)
}

