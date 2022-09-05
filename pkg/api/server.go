package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	mode string
	port string

	cors *cors.Config
}

// NewServer returns a new Server instance.
func NewServer(mode, port string, cors *cors.Config) *Server {
	return &Server{
		mode: mode,
		port: port,
		cors: cors,
	}
}

// Run is the Server runnable.
func (s *Server) Run() {
	gin.SetMode(s.mode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.New(*s.cors))

	// This grouping/versioning should match
	// the Cellosaurus data version.
	api := router.Group("/v41")

	api.GET("/", ListStatistics)
	api.GET("/cells", ListCells)
	api.GET("/cell-lines", ListCells)
	api.GET("/cell_lines", ListCells)
	api.GET("/cells/*id", FindCell)
	api.GET("/cell-lines/*id", FindCell)
	api.GET("/cell_lines/*id", FindCell)
	api.GET("/refs", ListReferences)
	api.GET("/xrefs", ListCrossReferences)
	api.GET("/stats", ListStatistics)

	router.NoRoute(BadRequest)

	router.Run(":" + s.port)
}
