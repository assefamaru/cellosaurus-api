package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	mode    string
	port    string
	version string

	cors *cors.Config
}

// NewServer returns a new Server instance.
func NewServer(mode, port, version string, cors *cors.Config) *Server {
	return &Server{
		mode:    mode,
		port:    port,
		version: version,
		cors:    cors,
	}
}

// Run is the Server runnable.
func (s *Server) Run() {
	gin.SetMode(s.mode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.New(*s.cors))
	router.Use(Logger())

	// This grouping/versioning matches
	// the Cellosaurus data version.
	api := router.Group("/api/v" + s.version)

	api.GET("/", ListStatistics)
	api.GET("/cells", ListCells)
	api.GET("/cell-lines", ListCells)
	api.GET("/cells/*id", FetchCell)
	api.GET("/cell-lines/*id", FetchCell)
	api.GET("/refs", ListReferences)
	api.GET("/references", ListReferences)
	api.GET("/xrefs", ListCrossReferences)
	api.GET("/cross-references", ListCrossReferences)
	api.GET("/stats", ListStatistics)
	api.GET("/statistics", ListStatistics)

	router.NoRoute(BadRequest)

	router.Run(":" + s.port)
}
