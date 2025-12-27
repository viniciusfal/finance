package repositories

import (
	"context"
	"fmt"

	"manager/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *entity.Category) error {
	query := `
		INSERT INTO categories (name, description, color, icon)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		category.Name,
		category.Description,
		category.Color,
		category.Icon,
	).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}

	return nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]entity.Category, error) {
	query := `
		SELECT id, name, description, color, icon, created_at, updated_at
		FROM categories
		ORDER BY name ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var cat entity.Category
		err := rows.Scan(
			&cat.ID,
			&cat.Name,
			&cat.Description,
			&cat.Color,
			&cat.Icon,
			&cat.CreatedAt,
			&cat.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int64) (*entity.Category, error) {
	query := `
		SELECT id, name, description, color, icon, created_at, updated_at
		FROM categories
		WHERE id = $1
	`

	var category entity.Category
	err := r.db.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.Color,
		&category.Icon,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	return &category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	query := `
		UPDATE categories
		SET name = $1, description = $2, color = $3, icon = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	err := r.db.QueryRow(ctx, query,
		category.Name,
		category.Description,
		category.Color,
		category.Icon,
		category.ID,
	).Scan(&category.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}

	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM categories WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

