package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"template-go-hexagonal/application/services"
	presentationerrors "template-go-hexagonal/presentation/errors"
	"template-go-hexagonal/presentation/responses"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.productService.ListCatalog(c.Request.Context())
	if err != nil {
		responses.ErrorFromApp(c, presentationerrors.Map(err))
		return
	}

	responses.Success(c, http.StatusOK, products)
}
