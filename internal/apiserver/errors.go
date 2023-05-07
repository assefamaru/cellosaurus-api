package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BadRequest(c *gin.Context) {
	errRenderer(c, http.StatusBadRequest, "Bad Request")
}

func errRenderer(c *gin.Context, code int, message string) {
	err := Error{
		Code:    code,
		Message: message,
	}
	c.JSON(code, gin.H{"error": err})
}
