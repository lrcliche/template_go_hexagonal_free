package tests

import (
	"context"
	"testing"

	"template-go-hexagonal/application/services"
	"template-go-hexagonal/domain/ports"

	"github.com/stretchr/testify/require"
)

func TestProductService_Update(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	created, err := service.Create(context.Background(), services.CreateProductInput{
		Name:        "Keyboard",
		Description: "Mechanical",
		Price:       99,
	})
	require.NoError(t, err)

	updated, err := service.Update(context.Background(), services.UpdateProductInput{
		ID:          created.ID,
		Name:        "Keyboard Pro",
		Description: "Mechanical RGB",
		Price:       149,
	})

	require.NoError(t, err)
	require.Equal(t, "Keyboard Pro", updated.Name)
	require.Equal(t, 149.0, updated.Price)
}

func TestProductService_Update_NotFound(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	updated, err := service.Update(context.Background(), services.UpdateProductInput{
		ID:          "404",
		Name:        "Keyboard Pro",
		Description: "Mechanical RGB",
		Price:       149,
	})

	require.Nil(t, updated)
	require.ErrorIs(t, err, ports.ErrProductNotFound)
}
