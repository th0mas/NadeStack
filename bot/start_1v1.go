package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/th0mas/NadeStack/csgo"
	"github.com/th0mas/NadeStack/models"
)

func (b *Bot) start1v1(s *discordgo.Session, m *discordgo.MessageCreate, cmd []string) {
	if len(cmd) < 3 {
		return
	}
	gameMap := cmd[1]
	if !(csgo.IsAvailiableMap(gameMap)) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is not a recognized map, aborting", gameMap))
		return
	}
	_, err := findUserVoiceChannel(s, m.GuildID, m.Author.ID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "User not currently in voice channel, panik")
	}
	p1, err := b.models.GetUserByDiscordID(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error getting player 1")
		return
	}

	if len(m.Mentions) != 1 {
		return
	}
	p2ID := m.Mentions[0].ID

	p2, err := b.models.GetUserByDiscordID(p2ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error getting player 2")
		return
	}
	teamOne := b.models.CreateTeam([]*models.User{
		p1,
	})
	teamTwo := b.models.CreateTeam([]*models.User{
		p2,
	})

	// TODO: Refractor out into common funcs

	s.ChannelMessageSend(m.ChannelID, "creating game")
	match := b.models.Create1v1(gameMap, *teamOne, *teamTwo)
	game := b.models.MakeGame(match)

	s.ChannelMessageSend(m.ChannelID, "creating embed")
	embed := initCsgoMatchEmbed(m.ChannelID, "1v1", game)
	embed.create(s)

	s.ChannelMessageSend(m.ChannelID, "running csgo service")
	csgo.BuildCSGOServer(game, b.conf.GSLT, func(g *models.Game) {
		embed.update(s)
	})

}
