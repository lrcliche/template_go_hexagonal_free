package container

import (
	"context"

	"template-go-hexagonal/application/config"
	"template-go-hexagonal/application/services"
	"template-go-hexagonal/domain/ports"
	"template-go-hexagonal/infrastructure/database"
	"template-go-hexagonal/infrastructure/repositories"
	"template-go-hexagonal/presentation/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	config            config.Config
	dbPool            *pgxpool.Pool
	productRepository ports.ProductRepository
	productService    services.ProductService
	productHandler    *handlers.ProductHandler
}

func NewContainer(ctx context.Context, cfg config.Config) (*Container, error) {
	dbPool, err := database.NewPostgresPool(ctx, cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	return &Container{config: cfg, dbPool: dbPool}, nil
}

func New(ctx context.Context, cfg config.Config) (*Container, error) {
	return NewContainer(ctx, cfg)
}

func (c *Container) Config() config.Config {
	return c.config
}

func (c *Container) Close() {
	if c.dbPool != nil {
		c.dbPool.Close()
	}
}

func (c *Container) ResolveProductRepository() ports.ProductRepository {
	if c.productRepository == nil {
		c.productRepository = repositories.NewPostgresProductRepository(c.dbPool)
	}
	return c.productRepository
}

func (c *Container) ResolveProductService() services.ProductService {
	if c.productService == nil {
		c.productService = services.NewProductService(c.ResolveProductRepository())
	}
	return c.productService
}

func (c *Container) ResolveProductHandler() *handlers.ProductHandler {
	if c.productHandler == nil {
		c.productHandler = handlers.NewProductHandler(c)
	}
	return c.productHandler
}
