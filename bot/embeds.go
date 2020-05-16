package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/th0mas/NadeStack/models"
)

const embedColour = 0xFF0000

type csgoMatchEmbed struct {
	gameType string
	game     *models.Game

	embedID string
	channelID string
	embed   *discordgo.MessageEmbed
}

func initCsgoMatchEmbed(channelID, gameType string, g *models.Game) *csgoMatchEmbed {
	c := &csgoMatchEmbed{
		gameType: gameType,
		channelID: channelID,
		game:     g,
	}

	return c
}

func (c *csgoMatchEmbed) buildFields() {
	c.embed = &discordgo.MessageEmbed{
		Title:       c.buildTitle(),
		Description: c.getDesc(),
		Timestamp:   "",
		Color:       0,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      createEmbedAuthor(),
		Fields:      nil,
	}
}

func (c *csgoMatchEmbed) buildTitle() string {
	verb := "Creating"
	if c.game.Status >= models.GameReady {
		verb =  "Playing"
	}

	return fmt.Sprintf("%s %s on %s", verb, c.gameType, c.game.Match.MapList[0])

}

func (c *csgoMatchEmbed) getDesc() string {
	// We return a description for whats gonna happen next as we switch on the current state of the server
	// e.g. status = Not started -> "Provisioning server..." as the bot is currently provisioning a server
	switch c.game.Status {
	case models.NotStarted:
		return "Provisioning Server..."
	case models.ServerProvisioned:
		return "Uploading plugins & info..."
	case models.ConfigUploaded:
		return "Unpacking config and info..."
	case models.ConfigUnpacked:
		return "Starting server..."
	case models.ServerStarted:
		return "Configuring server...."
	case models.ServerConfigured:
		return fmt.Sprintf("`connect %s; password nadestack`", *c.game.ServerIP)
	}

	return "No recognised status"
}

func (c *csgoMatchEmbed) create(s *discordgo.Session) {
	c.buildFields()
	s.ChannelMessageSendEmbed(c.channelID, c.embed)
}

func (c *csgoMatchEmbed) update(s *discordgo.Session) {
	c.buildFields()
	s.ChannelMessageEditEmbed(c.channelID, c.embedID, c.embed)
}

func createEmbedAuthor(s ...string) *discordgo.MessageEmbedAuthor {
	name := "Nadestack"
	if len(s) > 0 {
		name += " - " + s[0]
	}
	return &discordgo.MessageEmbedAuthor{
		IconURL: "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/320/twitter/154/firecracker_1f9e8.png",
		Name:    name,
	}
}

func createLinkEmbed(title, description, url string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author:      createEmbedAuthor(),
		Color:       embedColour,
		Description: description,
		//Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/320/twitter/154/firecracker_1f9e8.png"},
		URL:   url,
		Title: title,
		// Footer: &discordgo.MessageEmbedFooter{
		// 	Text: "NadeStack - create CSGO games on discord",
		// },
	}
}

func createUserInfoEmbed(u *models.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: createEmbedAuthor(),
		Color:  embedColour,
		Title:  u.DiscordNickname + "'s info",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "SteamID64",
				Value: *u.SteamID,
			},
			{
				Name:  "Date Created/Updated",
				Value: u.UpdatedAt.String(),
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: u.DiscordProfilePicURL},
	}
}
