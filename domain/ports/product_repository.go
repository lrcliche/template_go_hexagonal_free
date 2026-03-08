package ports

import (
	"context"

	"template-go-hexagonal/domain/entities"
)

// ProductRepository is a sample output port kept in the free version to show
// how application services depend on abstractions.
type ProductRepository interface {
	List(ctx context.Context) ([]entities.Product, error)
}
