package repositories

import (
	"context"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
)

// DemoProductRepository is intentionally non-functional in the free version.
// It exists only to demonstrate the adapter shape in hexagonal architecture.
type DemoProductRepository struct{}

func NewDemoProductRepository() *DemoProductRepository {
	return &DemoProductRepository{}
}

func (r *DemoProductRepository) List(ctx context.Context) ([]entities.Product, error) {
	_ = ctx
	return nil, ports.ErrNotImplemented
}
