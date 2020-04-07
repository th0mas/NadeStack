package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func GETSteamAuthURL(ctx *gin.Context, c *config.Config) {
	ctx.JSON(200, gin.H{
		"url: ": generateSteamOpenIdUrl(c).String(),
	})
}