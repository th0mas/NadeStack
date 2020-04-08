package web

import (
	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func getSteamCallback(ctx *gin.Context, c *config.Config) {
	steamID, err := verifySteamCallback(ctx, c)

	if err != nil {
		panic(err)
	}
	ctx.String(200, "Successfully validated steam id: "+steamID)
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
