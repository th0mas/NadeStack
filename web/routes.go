package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func getSteamCallback(ctx *gin.Context, c *config.Config) {
	steamID, err := verifySteamCallback(ctx, c)

	if err != nil {
		ctx.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
	ctx.JSON(200, gin.H{
		"success": true,
		"steamID": steamID,
	})}
}

func (w *Web) getDeeplinkInfo(ctx *gin.Context) {
	dl, err := w.model.GetDeepLinkData(ctx.Query("rune"))
	if w.model.IsRecordNotFound(err) {
		ctx.JSON(404, gin.H{
			"error": "not found",
		})
	} else {
		dl.Payload = generateSteamOpenIdUrl(w.conf).String()
		ctx.JSON(200, dl)
	}
}
