package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/user-service/log"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		status := c.Writer.Status()
		logger.Log.Infow("HTTP Request",
			"status", status,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"latency", latency,
			"client_ip", c.ClientIP(),
		)
	}
}
