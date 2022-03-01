package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Basic(c *gin.Context) {
	type Obj struct {
		Val string `json:"val"`
	}
	var obj Obj
	obj.Val = "basic"
	var err error
	logSentry(err)
	c.JSON(http.StatusOK, obj)
}
