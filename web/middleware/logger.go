package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware is a middleware function that logs incoming requests and responses.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// Process the request
		c.Next()

		// Calculate request duration
		duration := time.Since(start)

		// Log the request details
		logrus.Infof(
			"Method: %s, URL: %s, Status: %d, Duration: %s",
			c.Request.Method,
			c.Request.URL.String(),
			c.Writer.Status(),
			duration,
		)
	}
}
