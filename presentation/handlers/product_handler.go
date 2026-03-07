package handlers

import (
	"net/http"

	"template-go-hexagonal/application/services"
	presentationerrors "template-go-hexagonal/presentation/errors"
	"template-go-hexagonal/presentation/logging"
	"template-go-hexagonal/presentation/responses"

	"github.com/gin-gonic/gin"
)

type ProductServiceResolver interface {
	ResolveProductService() services.ProductService
}

type ProductHandler struct {
	serviceResolver ProductServiceResolver
}

type createProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

type updateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

func NewProductHandler(serviceResolver ProductServiceResolver) *ProductHandler {
	return &ProductHandler{serviceResolver: serviceResolver}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var request createProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		responses.ErrorFromApp(c, presentationerrors.BadRequest("Invalid request body"))
		return
	}

	product, err := h.serviceResolver.ResolveProductService().Create(c.Request.Context(), services.CreateProductInput{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	})
	if err != nil {
		h.handleError(c, "Create", err)
		return
	}

	responses.Success(c, http.StatusCreated, product)
}

func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.serviceResolver.ResolveProductService().List(c.Request.Context())
	if err != nil {
		h.handleError(c, "List", err)
		return
	}

	responses.Success(c, http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	product, err := h.serviceResolver.ResolveProductService().GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		h.handleError(c, "GetByID", err)
		return
	}

	responses.Success(c, http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	var request updateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		responses.ErrorFromApp(c, presentationerrors.BadRequest("Invalid request body"))
		return
	}

	product, err := h.serviceResolver.ResolveProductService().Update(c.Request.Context(), services.UpdateProductInput{
		ID:          c.Param("id"),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	})
	if err != nil {
		h.handleError(c, "Update", err)
		return
	}

	responses.Success(c, http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	if err := h.serviceResolver.ResolveProductService().Delete(c.Request.Context(), c.Param("id")); err != nil {
		h.handleError(c, "Delete", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProductHandler) handleError(c *gin.Context, operation string, err error) {
	logging.LogError("PRODUCT_HANDLER", operation, err)
	responses.ErrorFromApp(c, presentationerrors.Map(err))
}
