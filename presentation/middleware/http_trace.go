package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"template-go-hexagonal/presentation/logging"
)

func HTTPTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := c.Writer.Status()
		if status == 0 {
			status = http.StatusOK
		}

		logging.LogComponent(
			"HTTP_TRACE",
			fmt.Sprintf("method=%s path=%s status=%d duration=%s", c.Request.Method, c.Request.URL.Path, status, time.Since(start)),
		)
	}
}
