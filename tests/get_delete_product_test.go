package tests

import (
	"context"
	"testing"

	"template-go-hexagonal/application/services"
	"template-go-hexagonal/domain/ports"

	"github.com/stretchr/testify/require"
)

func TestProductService_GetByID(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	created, err := service.Create(context.Background(), services.CreateProductInput{Name: "Desk", Description: "Wood", Price: 40})
	require.NoError(t, err)

	found, err := service.GetByID(context.Background(), created.ID)

	require.NoError(t, err)
	require.Equal(t, created.ID, found.ID)
}

func TestProductService_Delete(t *testing.T) {
	repository := newFakeProductRepository()
	service := services.NewProductService(repository)

	created, err := service.Create(context.Background(), services.CreateProductInput{Name: "Desk", Description: "Wood", Price: 40})
	require.NoError(t, err)

	err = service.Delete(context.Background(), created.ID)
	require.NoError(t, err)

	deleted, err := service.GetByID(context.Background(), created.ID)
	require.Nil(t, deleted)
	require.ErrorIs(t, err, ports.ErrProductNotFound)
}
