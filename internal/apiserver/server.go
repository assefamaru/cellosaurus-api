package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/assefamaru/cellosaurus-api/internal/data"
	"github.com/assefamaru/cellosaurus-api/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	// The store API for accessing
	// ingested Cellosaurus data.
	store *data.ObjectStore
	// The server address.
	addr string
	// The server runtime mode.
	mode string
	// The Cellosaurus data version.
	version string
	// The server CORS settings.
	cors cors.Config
}

// New creates a new Server.
func New(ctx context.Context, client *db.MySQLClient, addr, mode, version string) *Server {
	return &Server{
		store:   data.NewObjectStore(client),
		addr:    addr,
		mode:    mode,
		version: version,
		cors:    defaultServerCORS(),
	}
}

func (s *Server) Run() {
	gin.SetMode(s.mode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.New(s.cors))
	router.Use(Logger())

	api := router.Group(s.version)
	api.GET("/", s.ListStatistics)
	api.GET("/cells", s.ListCells)
	api.GET("/cell-lines", s.ListCells)
	api.GET("/cell_lines", s.ListCells)
	api.GET("/cells/*id", s.GetCell)
	api.GET("/cell-lines/*id", s.GetCell)
	api.GET("/cell_lines/*id", s.GetCell)
	api.GET("/refs", s.ListReferences)
	api.GET("/references", s.ListReferences)
	api.GET("/xrefs", s.ListCrossReferences)
	api.GET("/cross-references", s.ListCrossReferences)
	api.GET("/cross_references", s.ListCrossReferences)
	api.GET("/stats", s.ListStatistics)
	api.GET("/statistics", s.ListStatistics)

	router.NoRoute(BadRequest)

	router.Run(s.addr)
}

func defaultServerCORS() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodOptions},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
