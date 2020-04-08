package bot

import (
	"fmt"
	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot implements the `Service` interface
type Bot struct {
	d  *discordgo.Session
	db *models.DB
}

// Run runs the discord-bot component of NadeStack
func (b *Bot) Run(c *config.Config, db *models.DB) {
	log.Println("Starting discord bot")
	b.db = db

	d, err := discordgo.New("Bot " + c.DiscordToken)

	if err != nil {
		panic(err)
	}

	d.AddHandler(b.messageCreateHandler)

	err = d.Open()
	if err != nil {
		panic(err)
	}

	b.d = d
}

func (b *Bot) Close() {
	_ = b.d.Close()
}

func (b *Bot) messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot - probaly not important
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!steamdebug") {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s: no linked steam account`", m.Author.Username))
	}

}
