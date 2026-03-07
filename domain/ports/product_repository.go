package ports

import (
	"context"

	"template-go-hexagonal/domain/entities"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) error
	List(ctx context.Context) ([]entities.Product, error)
	GetByID(ctx context.Context, id string) (*entities.Product, error)
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id string) error
}
