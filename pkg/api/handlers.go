package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Renders responses applying indentation settings.
func indentRenderer(c *gin.Context, data interface{}, indent bool) {
	if indent {
		c.IndentedJSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// GET /cell-lines.
func ListCells(c *gin.Context) {
	meta, err := getMeta(c, "cells")
	if err != nil {
		InternalServerError(c)
		return
	}
	indent := getIndent(c)
	cells := Cells{Meta: meta}
	if err = cells.List(); err != nil {
		InternalServerError(c)
		return
	}
	indentRenderer(c, cells, indent)
}

// GET /cell-lines/:id.
func FindCell(c *gin.Context) {
	indent := getIndent(c)
	cell := Cell{Identifier: strings.TrimPrefix(c.Param("id"), "/")}
	err := cell.Find()
	if err != nil {
		if err == sql.ErrNoRows {
			NotFound(c)
		} else {
			InternalServerError(c)
		}
		return
	}
	indentRenderer(c, cell, indent)
}

// GET /references.
func ListReferences(c *gin.Context) {
	meta, err := getMeta(c, "refs")
	if err != nil {
		InternalServerError(c)
		return
	}
	indent := getIndent(c)
	refs := References{Meta: meta}
	if err = refs.List(); err != nil {
		InternalServerError(c)
		return
	}
	indentRenderer(c, refs, indent)
}
