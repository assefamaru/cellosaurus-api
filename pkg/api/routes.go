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
	Route{http.MethodGet, "/", Basic},
}
