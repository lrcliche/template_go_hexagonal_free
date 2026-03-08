package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"template-go-hexagonal/presentation/container"
	"template-go-hexagonal/presentation/middleware"
	"template-go-hexagonal/presentation/responses"
)

func New(appContainer *container.Container) *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.Recovery())

	engine.GET("/health", func(c *gin.Context) {
		responses.Success(c, http.StatusOK, gin.H{"status": "ok"})
	})

	productHandler := appContainer.ProductHandler()

	v1 := engine.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", productHandler.List)
		}
	}

	return engine
}
