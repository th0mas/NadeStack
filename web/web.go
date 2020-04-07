package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func Run(c *config.Config) {
	r := gin.Default()

	r.GET("/health", healthCheck)

	r.GET("/steamurl", func(context *gin.Context) {
		GETSteamAuthURL(context, c)
	})

	_ = r.Run()
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}