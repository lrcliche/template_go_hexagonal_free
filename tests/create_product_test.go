package tests

import (
	"context"
	"testing"

	"template-go-hexagonal/application/services"
	"template-go-hexagonal/domain/entities"

	"github.com/stretchr/testify/require"
)

func TestProductService_Create(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	product, err := service.Create(context.Background(), services.CreateProductInput{
		Name:        "Laptop",
		Description: "Business laptop",
		Price:       1000,
	})

	require.NoError(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, "Laptop", product.Name)
}

func TestProductService_Create_InvalidPrice(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	product, err := service.Create(context.Background(), services.CreateProductInput{
		Name:        "Laptop",
		Description: "Business laptop",
		Price:       0,
	})

	require.Nil(t, product)
	require.ErrorIs(t, err, entities.ErrInvalidProductPrice)
}
