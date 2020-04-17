package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func (w *Web) getSteamCallback(ctx *gin.Context, c *config.Config) {
	actionID := ctx.Query("rune")
	steamID, err := verifySteamCallback(ctx, c, actionID)

	if err != nil {
		ctx.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"steamID": steamID,
		})
	}
}

func (w *Web) getDeeplinkInfo(ctx *gin.Context) {
	dl, err := w.model.GetDeepLinkData(ctx.Query("rune"))
	if w.model.IsRecordNotFound(err) {
		ctx.JSON(404, gin.H{
			"error": "not found",
		})
	} else {
		dl.Payload = generateSteamOpenIdUrl(w.conf, ctx.Query("rune")).String()
		ctx.JSON(200, dl)
	}
}
