package bot

import (
	"fmt"
	"github.com/th0mas/NadeStack/config"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Run runs the discord-bot component of NadeStack
func Run(c *config.Config) *discordgo.Session {
	log.Println("Starting discord bot")

	d, err := discordgo.New("Bot " + c.DiscordToken)

	if err != nil {
		panic(err)
	}

	d.AddHandler(messageCreateHandler)

	err = d.Open()
	if err != nil {
		panic(err)
	}

	return d
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot - probaly not important
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!steamdebug") {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s: no linked steam account`", m.Author.Username))
	}

}
