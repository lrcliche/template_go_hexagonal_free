package responses

import (
	"github.com/gin-gonic/gin"
	presentationerrors "template-go-hexagonal/presentation/errors"
)

type Response struct {
	Errors []ErrorItem `json:"errors"`
	Data   any         `json:"data"`
}

type ErrorResponse struct {
	Errors []ErrorItem `json:"errors"`
}

type ErrorItem struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Success(c *gin.Context, statusCode int, payload any) {
	c.JSON(statusCode, Response{Errors: []ErrorItem{}, Data: payload})
}

func Error(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, ErrorResponse{Errors: []ErrorItem{{Code: code, Message: message}}})
}

func ErrorFromApp(c *gin.Context, appError presentationerrors.AppError) {
	Error(c, appError.Status, appError.Code, appError.Message)
}

func Forbidden(c *gin.Context) {
	c.Status(403)
}
