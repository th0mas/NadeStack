package web

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"
)

// Web implements the `Service` interface
type Web struct {
	r     *gin.Engine
	model *models.Models
	conf  *config.Config
}

// Run runs the web service component
func (w *Web) Run(c *config.Config, db *models.Models) {
	w.model = db
	w.conf = c
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./web/nadestack-frontend/build/", false)))
	api := r.Group("/api")

	api.GET("/health", healthCheck)
	api.GET("/deeplink", w.getDeeplinkInfo)
	api.GET("/auth/callback", func(context *gin.Context) { // TODO: Doesn't need extra func
		w.getSteamCallback(context, c)
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./web/nadestack-frontend/build/index.html")
	})

	w.r = r
	_ = r.Run()
}

// Close is a dummy method as we do not need to explicitly close the web server
func (w *Web) Close() {}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}
