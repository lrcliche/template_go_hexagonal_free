package errors

import (
	"errors"
	"net/http"

	"template-go-hexagonal/domain/ports"
)

func Map(err error) AppError {
	switch {
	case errors.Is(err, ports.ErrNotImplemented):
		return New(http.StatusNotImplemented, "not-implemented", "This endpoint is a demo in the free version")
	default:
		return New(http.StatusInternalServerError, "internal-server-error", "Internal Error")
	}
}

func BadRequest(message string) AppError {
	return New(http.StatusBadRequest, "bad-request", message)
}
