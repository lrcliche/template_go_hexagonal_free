package tests

import (
	"context"
	"strconv"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
)

type fakeProductRepository struct {
	items  map[string]*entities.Product
	nextID int
}

func newFakeProductRepository() *fakeProductRepository {
	return &fakeProductRepository{
		items:  make(map[string]*entities.Product),
		nextID: 1,
	}
}

func (r *fakeProductRepository) Create(_ context.Context, product *entities.Product) error {
	product.ID = strconv.Itoa(r.nextID)
	r.nextID++
	copyValue := *product
	r.items[product.ID] = &copyValue
	return nil
}

func (r *fakeProductRepository) List(_ context.Context) ([]entities.Product, error) {
	products := make([]entities.Product, 0, len(r.items))
	for _, product := range r.items {
		products = append(products, *product)
	}
	return products, nil
}

func (r *fakeProductRepository) GetByID(_ context.Context, id string) (*entities.Product, error) {
	product, ok := r.items[id]
	if !ok {
		return nil, ports.ErrProductNotFound
	}
	copyValue := *product
	return &copyValue, nil
}

func (r *fakeProductRepository) Update(_ context.Context, product *entities.Product) error {
	if _, ok := r.items[product.ID]; !ok {
		return ports.ErrProductNotFound
	}
	copyValue := *product
	r.items[product.ID] = &copyValue
	return nil
}

func (r *fakeProductRepository) Delete(_ context.Context, id string) error {
	if _, ok := r.items[id]; !ok {
		return ports.ErrProductNotFound
	}
	delete(r.items, id)
	return nil
}
