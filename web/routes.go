package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func getSteamAuthURL(ctx *gin.Context, c *config.Config) {
	ctx.JSON(200, gin.H{
		"url: ": generateSteamOpenIdUrl(c).String(),
	})
}

func getSteamCallback(ctx *gin.Context, c *config.Config) {
	steamID, err := verifySteamCallback(ctx, c)

	if err != nil {
		panic(err)
	}
	ctx.String(200, "Successfully validated steam id: "+steamID)
}
