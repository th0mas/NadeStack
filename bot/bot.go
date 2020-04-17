package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"

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
		b.steamDebugCommand(s, m)
	}

	if strings.HasPrefix(m.Content, "/linksteam") {
		b.steamLinkCommand(s, m)
	}

}

func (b *Bot) steamLinkCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u, err := b.models.GetUserByDiscordID(m.Author.ID)
	if b.models.IsRecordNotFound(err) {
		if b.conf.Debug {
			_, _ = s.ChannelMessageSend(m.ChannelID, "`DEBUG: No user with Discord ID "+m.Author.ID+" creating user")
		}
		u = b.models.CreateUserFromDiscord(m.Author.ID, m.Author.Username, m.Author.Avatar)
	} else {
		if b.conf.Debug {
			_, _ = s.ChannelMessageSend(m.ChannelID, "`DEBUG: User already exists, skipping create`")
		}
	}
	userChannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		panic(err)
	}

	dl := b.models.CreateDeepLink(models.LinkSteamID, u)
	if b.conf.Debug {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`DEBUG: Created Deep Link with vals %+v`", dl))
	}
	_, _ = s.ChannelMessageSend(userChannel.ID, fmt.Sprintf("Click the link to link you accounts: %s/%s", b.conf.Domain, dl.ShortURL))

	_ = s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
}

func (b *Bot) steamDebugCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u, err := b.models.GetUserByDiscordID(m.Author.ID)
	if err != nil || u.SteamID == nil {
		s.ChannelMessageSend(m.ChannelID, "no user infomation found for discord id: "+m.Author.ID)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Steam connection infomation: \n SteamID: `%s` \n SteamID64: '%d'",
		*u.SteamID, u.SteamID64))
}
