package logger

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger - Logs info about incoming requests
func Logger(inner func(*gin.Context), name string) func(*gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()

		inner(c)

		log.Printf(
			"%s\t%s\t%s\t%s",
			c.Request.Method,
			c.Request.RequestURI,
			name,
			time.Since(start),
		)
	}
}
