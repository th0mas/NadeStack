package web

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

func (w *Web) getSteamCallback(ctx *gin.Context, c *config.Config) {
	actionID := ctx.Query("rune")
	steamID64Str, err := verifySteamCallback(ctx, c, actionID)

	if err != nil {
		ctx.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}

	linkData, err := w.model.GetDeepLinkData(actionID)

	if err != nil {
		ctx.JSON(500, gin.H{"success": false,
			"error": "failed to get client, did you refresh the page?",
		})
		return
	}

	steamID64, err := strconv.ParseUint(steamID64Str, 10, 64)

	if err != nil {
		ctx.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}

	steamID := ConvertSteamID64toSteamID3(steamID64)

	w.model.AddSteamIDToUser(&linkData.User, steamID, steamID64)

	ctx.JSON(200, gin.H{
		"success":   true,
		"steamID64": steamID64,
		"steamID":   steamID,
		"nickname":  linkData.User.DiscordNickname,
	})
}

func (w *Web) getDeeplinkInfo(ctx *gin.Context) {
	dl, err := w.model.GetDeepLinkData(ctx.Query("rune"))
	if w.model.IsRecordNotFound(err) {
		ctx.JSON(404, gin.H{
			"error": "not found",
		})
	} else {
		dl.Payload = generateSteamOpenIDUrl(w.conf, ctx.Query("rune")).String()
		ctx.JSON(200, dl)
	}
}

func (w *Web) getMatchInfo(ctx *gin.Context) {
	matchID := ctx.Query("id")

	m, err := w.model.GetMatchByID(matchID)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "could not find match"})
		return
	}
	m.GenerateTeamIDS()
	ctx.JSON(200, m)
	return
}
