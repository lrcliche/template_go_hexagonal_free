package services

import (
	"context"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
)

// ProductService demonstrates the application layer contract.
type ProductService interface {
	ListCatalog(ctx context.Context) ([]entities.Product, error)
}

type productService struct {
	productRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) ListCatalog(ctx context.Context) ([]entities.Product, error) {
	// TODO(premium): apply filters, pagination, and business rules.
	return s.productRepository.List(ctx)
}
