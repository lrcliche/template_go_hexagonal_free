package routes

import (
	"net/http"

	"template-go-hexagonal/presentation/container"
	"template-go-hexagonal/presentation/middleware"
	"template-go-hexagonal/presentation/responses"

	"github.com/gin-gonic/gin"
)

func New(appContainer *container.Container) *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.HTTPTrace())
	engine.Use(middleware.Recovery())

	engine.GET("/health", func(c *gin.Context) {
		responses.Success(c, http.StatusOK, gin.H{"status": "ok"})
	})

	productHandler := appContainer.ResolveProductHandler()

	v1 := engine.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", productHandler.Create)
			products.GET("", productHandler.List)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}
	}

	return engine
}
