package bot

import "github.com/bwmarrin/discordgo"

func createEmbed(title, description, url string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/320/twitter/154/firecracker_1f9e8.png",
			Name:    "NadeStack",
		},
		Color:       0xFF0000,
		Description: description,
		//Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/320/twitter/154/firecracker_1f9e8.png"},
		URL:   url,
		Title: title,
		// Footer: &discordgo.MessageEmbedFooter{
		// 	Text: "NadeStack - create CSGO games on discord",
		// },
	}
}
