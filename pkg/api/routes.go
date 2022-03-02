package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method   string
	Endpoint string
	Handler  gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{http.MethodGet, "/cells", ListCells},
	Route{http.MethodGet, "/cell_lines", ListCells},
	Route{http.MethodGet, "/cell-lines", ListCells},
	Route{http.MethodGet, "/cells/*id", FindCell},
	Route{http.MethodGet, "/cell_lines/*id", FindCell},
	Route{http.MethodGet, "/cell-lines/*id", FindCell},
	Route{http.MethodGet, "/refs", ListReferences},
	Route{http.MethodGet, "/xrefs", ListCrossReferences},
}
