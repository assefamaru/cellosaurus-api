package api

import "github.com/gin-gonic/gin"

type Context struct {
	Mode string
	Port string
}

func Init(ctx Context) {
	// Set db credentials
	SetupMySQLConfig()

	// Set gin mode
	gin.SetMode(ctx.Mode)

	// Gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// This grouping corresponds to
	// the cellosaurus data version
	v := router.Group("/v41")
	for _, route := range routes {
		v.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match request URL,
	// return 400 (Bad Request)
	router.NoRoute(BadRequest)

	// Listen and serve on 0.0.0.0:PORT
	router.Run(":" + ctx.Port)
}
