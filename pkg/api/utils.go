package api

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMeta(c *gin.Context, resource string) (Meta, error) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	// Set max number of cell lines to return per request
	if perPage > 100 {
		perPage = 100
	}

	switch resource {
	case "refs":
		total, err := totalRefs()
		if err != nil {
			return Meta{}, err
		}
		lastPage := int(math.Ceil(float64(total) / float64(perPage)))
		meta := Meta{page, perPage, lastPage, total}
		return meta, nil
	default:
		total, err := totalCells()
		if err != nil {
			return Meta{}, err
		}
		lastPage := int(math.Ceil(float64(total) / float64(perPage)))
		meta := Meta{page, perPage, lastPage, total}
		return meta, nil
	}
}

func getIndent(c *gin.Context) bool {
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))
	return indent
}

func getPaginationFrom(meta Meta) int {
	return (meta.Page - 1) * meta.PerPage
}
