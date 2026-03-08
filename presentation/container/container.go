package container

import (
	"template-go-hexagonal/application/services"
	"template-go-hexagonal/infrastructure/repositories"
	"template-go-hexagonal/presentation/handlers"
)

// Container in the free version keeps wiring explicit and simple.
type Container struct {
	productHandler *handlers.ProductHandler
}

func New() *Container {
	productRepository := repositories.NewDemoProductRepository()
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	return &Container{productHandler: productHandler}
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}
