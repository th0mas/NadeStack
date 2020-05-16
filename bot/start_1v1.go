package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

	match := b.models.Create1v1(gameMap, *teamOne, *teamTwo)
	game := b.models.MakeGame(match)

	embed := initCsgoMatchEmbed(m.ChannelID, "1v1", game)
	embed.create(s)

	serverID, _ := csgo.CreateCSGOServer(5, gameMap, match.ID, b.conf.GSLT)
	game.IncrementGameStatus()
	game.ServerID = &serverID
	embed.update(s)

	game.IncrementGameStatus()
	embed.update(s)
	csgo.UploadPluginsToServer(serverID)

	match.GenerateTeamIDS()
	conf, _ := json.Marshal(match)
	csgo.UploadJSONToServer(serverID, "config.json", conf, "file")

	game.IncrementGameStatus()
	embed.update(s)
	csgo.UnzipPlugins(serverID)

	game.IncrementGameStatus()
	embed.update(s)
	csgo.StartServer(serverID)

	serverIP, err := waitForServerIP(serverID)

	game.ServerIP = &serverIP
	game.IncrementGameStatus()
	embed.update(s)

	err = csgo.SendCommandToServer(serverID, "get5_loadmatch config.json")

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to config server")
	}

	game.IncrementGameStatus()
	embed.update(s)

}

func waitForServerIP(id string) (string, error) {
	ch := make(chan string, 1)
	defer close(ch)

	go func() {
		for {
			if status, ip := csgo.GetServerStatus(id); !status {
				ch <- ip
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	timer := time.NewTimer(2 * time.Minute)
	defer timer.Stop()

	select {
	case ip := <-ch:
		return ip, nil
	case <-timer.C:
		return "", errors.New("timed out waiting for serve to load")
	}

}
