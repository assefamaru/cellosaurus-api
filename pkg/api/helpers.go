package api

import (
	"math"
	"strconv"

	"github.com/assefamaru/cellosaurus-api/pkg/api/core"
	"github.com/gin-gonic/gin"
)

const (
	maxResponseBatchSize = 1000
)

// Metadata for paginated results.
type Meta struct {
	Page     int `json:"page"`
	PerPage  int `json:"perPage"`
	LastPage int `json:"lastPage"`
	Total    int `json:"total"`
}

func newMeta(c *gin.Context, resourceType string) (*Meta, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("perPage", "10"))
	if err != nil {
		return nil, err
	}

	// Set hard upper limit on number of records in response.
	if perPage > maxResponseBatchSize {
		perPage = maxResponseBatchSize
	}

	var count int
	switch resourceType {
	case "references":
		count, err = core.CountReferences()
		if err != nil {
			return nil, err
		}
	default:
		count, err = core.CountCells()
		if err != nil {
			return nil, err
		}
	}
	lastPage := int(math.Ceil(float64(count) / float64(perPage)))
	meta := &Meta{page, perPage, lastPage, count}

	return meta, nil
}

func paginationFrom(meta *Meta) int {
	return (meta.Page - 1) * meta.PerPage
}
