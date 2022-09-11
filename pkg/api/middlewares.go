package api

import (
	"strings"
	"time"

	"github.com/assefamaru/cellosaurus-api/pkg/logging"
	"github.com/gin-gonic/gin"
)

// Logger is a custom middleware for logging requests.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		path := "/" + strings.TrimPrefix(c.Request.URL.Path, "/")
		logFormat := "%6s %s (%v) - %s"

		if c.Writer.Status() < 400 {
			logging.Infof(logFormat, c.Request.Method, path, latency, c.ClientIP())
		} else {
			logging.Errorf(logFormat, c.Request.Method, path, latency, c.ClientIP())
		}
	}
}
