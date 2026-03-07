package services

import (
	"context"
	"log"
	"time"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
)

type CreateProductInput struct {
	Name        string
	Description string
	Price       float64
}

type UpdateProductInput struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

type ProductService interface {
	Create(ctx context.Context, input CreateProductInput) (*entities.Product, error)
	Update(ctx context.Context, input UpdateProductInput) (*entities.Product, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*entities.Product, error)
	List(ctx context.Context) ([]entities.Product, error)
}

type productService struct {
	productRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) Create(ctx context.Context, input CreateProductInput) (*entities.Product, error) {
	now := time.Now().UTC()
	product := &entities.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := product.Validate(); err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=Create error=%v", err)
		return nil, err
	}

	if err := s.productRepository.Create(ctx, product); err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=Create error=%v", err)
		return nil, err
	}

	return product, nil
}

func (s *productService) List(ctx context.Context) ([]entities.Product, error) {
	products, err := s.productRepository.List(ctx)
	if err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=List error=%v", err)
		return nil, err
	}

	return products, nil
}

func (s *productService) GetByID(ctx context.Context, id string) (*entities.Product, error) {
	product, err := s.productRepository.GetByID(ctx, id)
	if err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=GetByID error=%v", err)
		return nil, err
	}

	return product, nil
}

func (s *productService) Update(ctx context.Context, input UpdateProductInput) (*entities.Product, error) {
	existingProduct, err := s.productRepository.GetByID(ctx, input.ID)
	if err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=Update error=%v", err)
		return nil, err
	}

	existingProduct.Name = input.Name
	existingProduct.Description = input.Description
	existingProduct.Price = input.Price
	existingProduct.UpdatedAt = time.Now().UTC()

	if err := existingProduct.Validate(); err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=Update error=%v", err)
		return nil, err
	}

	if err := s.productRepository.Update(ctx, existingProduct); err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=Update error=%v", err)
		return nil, err
	}

	return existingProduct, nil
}

func (s *productService) Delete(ctx context.Context, id string) error {
	if err := s.productRepository.Delete(ctx, id); err != nil {
		log.Printf("[PRODUCT_SERVICE] operation=Delete error=%v", err)
		return err
	}

	return nil
}
