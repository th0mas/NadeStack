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
	d      *discordgo.Session
	models *models.Models
	conf   *config.Config
}

// Run runs the discord-bot component of NadeStack
func (b *Bot) Run(c *config.Config, db *models.Models) {
	log.Println("Starting discord bot")
	b.models = db
	b.conf = c

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

	if strings.HasPrefix(m.Content, "/steamdebug") {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s: no linked steam account`", m.Author.Username))
	}

	if strings.HasPrefix(m.Content, "/linksteam") {
		b.steamLinkCommand(s, m)
	}

}

func (b *Bot) steamLinkCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, e := b.models.GetUserByDiscordID(m.Author.ID); b.models.CheckUserNotFound(e) {
		if b.conf.Debug {
			_, _ = s.ChannelMessageSend(m.ChannelID, "`DEBUG: No user with Discord ID "+m.Author.ID+" creating user")
		}
		b.models.CreateUserFromDiscord(m.Author.ID, m.Author.Username, m.Author.Avatar)
	} else {
		if b.conf.Debug {
			_, _ = s.ChannelMessageSend(m.ChannelID, "`DEBUG: User already exists, skipping create`")
		}
	}
	userChannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		panic(err)
	}
	_, _ = s.ChannelMessageSend(userChannel.ID, fmt.Sprintf("Yes this should be steam link"))
}
