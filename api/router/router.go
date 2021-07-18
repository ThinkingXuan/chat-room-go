package router

import (
	"github.com/gin-gonic/gin"
)

// Load load the middlewares, routers
func Load(e *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// use middlewares
	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(mw...)

	return e
}
