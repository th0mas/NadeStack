package bot

import (
	"log"
	"strings"

	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"

	"github.com/bwmarrin/discordgo"
)

const cmdPrefix = "/"

// Bot implements the `Service` interface
type Bot struct {
	d      *discordgo.Session
	models *models.Models
	conf   *config.Config

	commands map[string]commandHandler
}

type commandHandler func(s *discordgo.Session, m *discordgo.MessageCreate)

// Run runs the discord-bot component of NadeStack
func (b *Bot) Run(c *config.Config, db *models.Models) {
	log.Println("Starting discord bot")
	b.models = db
	b.conf = c
	b.commands = make(map[string]commandHandler)

	d, err := discordgo.New("Bot " + c.DiscordToken)

	if err != nil {
		panic(err)
	}

	// Register our own commands here
	b.addCommand("steamdebug", b.steamDebugCommand)
	b.addCommand("linksteam", b.steamLinkCommand)

	// Register a message handler with the discord API
	d.AddHandler(b.messageCreateHandler)

	err = d.Open()
	if err != nil {
		panic(err)
	}

	b.d = d
}

// Close closes the bots connection to discord
func (b *Bot) Close() {
	_ = b.d.Close()
}

func (b *Bot) addCommand(command string, handler commandHandler) {
	// TODO: test for existing command
	b.commands[command] = handler
}

func (b *Bot) messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot - probaly not important
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd := strings.Split(strings.TrimSpace(m.Content), " ")

	if !(strings.HasPrefix(cmd[0], cmdPrefix)) {
		return
	}

	fn, exists := b.commands[cmd[0][1:]]

	if exists {
		fn(s, m)
	}

}
