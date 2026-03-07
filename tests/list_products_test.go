package tests

import (
	"context"
	"testing"

	"template-go-hexagonal/application/services"

	"github.com/stretchr/testify/require"
)

func TestProductService_List(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	_, _ = service.Create(context.Background(), services.CreateProductInput{Name: "Mouse", Description: "Wireless", Price: 10})
	_, _ = service.Create(context.Background(), services.CreateProductInput{Name: "Monitor", Description: "27 inch", Price: 20})

	products, err := service.List(context.Background())

	require.NoError(t, err)
	require.Len(t, products, 2)
}
