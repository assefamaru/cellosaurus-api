package api

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"
	"strings"

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

// GET /cells.
func ListCells(c *gin.Context) {
	meta, err := newMeta(c, "cells")
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	cells, err := core.ListCells(paginationFrom(meta), meta.PerPage)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"meta": meta, "data": cells})
}

// GET /cells/:id.
func FetchCell(c *gin.Context) {
	cell, err := core.FetchCell(strings.TrimPrefix(c.Param("id"), "/"))
	if err == sql.ErrNoRows {
		errRenderer(c, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cell)
}

// GET /refs.
func ListReferences(c *gin.Context) {
	meta, err := newMeta(c, "refs")
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	refs, err := core.ListReferences(paginationFrom(meta), meta.PerPage)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"meta": meta, "data": refs})
}

// GET /xrefs.
func ListCrossReferences(c *gin.Context) {
	xrefs, err := core.ListCrossReferences()
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, xrefs)
}

// GET /stats.
func ListStatistics(c *gin.Context) {
	var version string
	pathParts := strings.Split(c.Request.URL.Path, "/")
	if len(pathParts) > 2 {
		version = strings.TrimPrefix(pathParts[2], "v")
	}

	stats, err := core.ListStatistics(version)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stats)
}

// newMeta returns metadata field with pagination information for list operations.
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
