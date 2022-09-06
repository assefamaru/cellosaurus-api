package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/assefamaru/cellosaurus-api/pkg/api/core"
	"github.com/gin-gonic/gin"
)

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
func FindCell(c *gin.Context) {
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
	xrefs, err := core.ListXRefs()
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, xrefs)
}

// GET /stats.
func ListStatistics(c *gin.Context) {
	stats, err := core.FetchStatistics()
	if err != nil {
		errRenderer(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stats)
}
