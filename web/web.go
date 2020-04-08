package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"
)

// Web implements the `Service` interface
type Web struct {
	r  *gin.Engine
	db *models.Models
}

// Run runs the web service component
func (w *Web) Run(c *config.Config, db *models.Models) {
	w.db = db
	r := gin.Default()

	r.GET("/health", healthCheck)

	r.GET("/steamurl", func(context *gin.Context) {
		getSteamAuthURL(context, c)
	})
	r.GET("/auth/callback", func(context *gin.Context) {
		getSteamCallback(context, c)
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
