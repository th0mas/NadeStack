package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/th0mas/NadeStack/models"
)

const embedColour = 0xFF0000

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

func createCSGOMatchEmbed(title, description, gameMap string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author:      createEmbedAuthor("5v5 on " + gameMap),
		Color:       embedColour,
		Title:       title,
		Description: description,

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Team One",
				Value:  "player 1 \n player 2 \n player 3 \n player 4 \n player 5\n",
				Inline: true,
			},
			{
				Name:   "Team Two",
				Value:  "player 1 \n player 2 \n player 3 \n player 4 \n player 5\n",
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://vignette.wikia.nocookie.net/cswikia/images/a/a7/CSGO_de_Mirage.jpg"},
		Footer:    &discordgo.MessageEmbedFooter{Text: "Game ID: aaaaaaa"},
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
