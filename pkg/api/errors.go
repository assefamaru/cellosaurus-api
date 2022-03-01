package api

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func logSentry(err error) {
	sentry.CaptureException(err)
}

func errRenderer(c *gin.Context, status int, message string, indent bool) {
	err := Error{
		Status:  status,
		Message: message,
	}
	if indent {
		c.IndentedJSON(status, gin.H{"error": err})
	} else {
		c.JSON(status, gin.H{"error": err})
	}
}

func BadRequest(c *gin.Context) {
	errRenderer(c, http.StatusBadRequest, "Bad Request", false)
}

func NotFound(c *gin.Context) {
	errRenderer(c, http.StatusNotFound, "Resource Not Found", false)
}

func InternalServerError(c *gin.Context) {
	errRenderer(c, http.StatusInternalServerError, "Internal Server Error", false)
}
