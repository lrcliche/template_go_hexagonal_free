package repositories

import (
	"context"
	"errors"
	"log"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
	domainports "template-go-hexagonal/domain/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresProductRepository struct {
	db *pgxpool.Pool
}

var _ domainports.ProductRepository = (*PostgresProductRepository)(nil)

func NewPostgresProductRepository(db *pgxpool.Pool) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(ctx context.Context, product *entities.Product) error {
	query := `
		INSERT INTO products (name, description, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	return r.db.QueryRow(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.ID)
}

func (r *PostgresProductRepository) List(ctx context.Context) ([]entities.Product, error) {
	query := `
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]entities.Product, 0)
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *PostgresProductRepository) GetByID(ctx context.Context, id string) (*entities.Product, error) {
	query := `
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	var product entities.Product
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ports.ErrProductNotFound
	}
	if err != nil {
		log.Printf("[POSTGRES_PRODUCT_REPOSITORY] operation=GetByID error=%v", err)
		return nil, err
	}

	return &product, nil
}

func (r *PostgresProductRepository) Update(ctx context.Context, product *entities.Product) error {
	query := `
		UPDATE products
		SET name = $2,
			description = $3,
			price = $4,
			updated_at = $5
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.UpdatedAt,
	)
	if err != nil {
		log.Printf("[POSTGRES_PRODUCT_REPOSITORY] operation=Update error=%v", err)
		return err
	}

	if result.RowsAffected() == 0 {
		return ports.ErrProductNotFound
	}

	return nil
}

func (r *PostgresProductRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.Exec(ctx, `DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		log.Printf("[POSTGRES_PRODUCT_REPOSITORY] operation=Delete error=%v", err)
		return err
	}

	if result.RowsAffected() == 0 {
		return ports.ErrProductNotFound
	}

	return nil
}
