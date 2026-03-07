package errors

import (
	"errors"
	"net/http"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
)

func Map(err error) AppError {
	switch {
	case errors.Is(err, ports.ErrProductNotFound):
		return New(http.StatusNotFound, "product-not-found", "Product not found")
	case errors.Is(err, entities.ErrInvalidProductName), errors.Is(err, entities.ErrInvalidProductPrice):
		return New(http.StatusUnprocessableEntity, "validation-error", err.Error())
	default:
		return New(http.StatusInternalServerError, "internal-server-error", "Internal Error")
	}
}

func BadRequest(message string) AppError {
	return New(http.StatusBadRequest, "bad-request", message)
}

func Conflict(code, message string) AppError {
	return New(http.StatusConflict, code, message)
}
