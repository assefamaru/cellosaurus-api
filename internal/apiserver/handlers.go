package apiserver

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/assefamaru/cellosaurus-api/internal/data"
	"github.com/gin-gonic/gin"
)

const (
	maxResponseBatchSize = 1000
)

// Metadata for paginated results.
type meta struct {
	Page     int `json:"page"`
	PerPage  int `json:"perPage"`
	LastPage int `json:"lastPage"`
	Total    int `json:"total"`
}

// GET /vX/cells.
func (s *Server) ListCells(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("perPage", "10")
	meta, err := s.newMeta(pageStr, perPageStr, data.ResourceTypeCell)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}
	cells, err := s.store.ListCells(paginationFrom(meta), meta.PerPage)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": meta, "data": cells})
}

// GET /vX/cells/:id.
func (s *Server) GetCell(c *gin.Context) {
	cell, err := s.store.FetchCell(strings.TrimPrefix(c.Param("id"), "/"))
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

// GET /vX/refs.
func (s *Server) ListReferences(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("perPage", "10")
	meta, err := s.newMeta(pageStr, perPageStr, data.ResourceTypeRef)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}
	refs, err := s.store.ListReferences(paginationFrom(meta), meta.PerPage)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": meta, "data": refs})
}

// GET /vX/xrefs.
func (s *Server) ListCrossReferences(c *gin.Context) {
	xrefs, err := s.store.ListCrossReferences()
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, xrefs)
}

// GET /vX/stats.
func (s *Server) ListStatistics(c *gin.Context) {
	var version string
	pathParts := strings.Split(c.Request.URL.Path, "/")
	if len(pathParts) > 2 {
		version = strings.TrimPrefix(pathParts[1], "v")
	}
	stats, err := s.store.ListStatistics(version)
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (s *Server) newMeta(pageStr, perPageStr, resourceType string) (*meta, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		return nil, err
	}
	// Set a ceiling on the allowed
	// number of records in response.
	if perPage > maxResponseBatchSize {
		perPage = maxResponseBatchSize
	}
	var count int
	switch resourceType {
	case data.ResourceTypeRef:
		count, err = s.store.CountReferences()
	case data.ResourceTypeCell:
		count, err = s.store.CountCells()
	}
	if err != nil {
		return nil, err
	}
	lastPage := int(math.Ceil(float64(count) / float64(perPage)))
	meta := &meta{page, perPage, lastPage, count}
	return meta, nil
}

func paginationFrom(meta *meta) int {
	return (meta.Page - 1) * meta.PerPage
}
