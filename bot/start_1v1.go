package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/th0mas/NadeStack/csgo"
	"github.com/th0mas/NadeStack/models"
	"time"
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

	match := b.models.Create1v1(gameMap, *teamOne, *teamTwo)

	game := b.models.MakeGame(match)
	fmt.Println(*game)
	embed := createCSGOMatchEmbed("Creating 1v1", "Provisioning server...", gameMap)
	status, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	serverID, _ := csgo.CreateCSGOServer(5, gameMap, match.ID, b.conf.GSLT)

	embed = createCSGOMatchEmbed("Creating 1v1", "Uploading info...", gameMap)
	s.ChannelMessageEditEmbed(m.ChannelID, status.ID, embed)
	csgo.UploadPluginsToServer(serverID)
	
	match.GenerateTeamIDS()
	conf, _ := json.Marshal(match)
	csgo.UploadJSONToServer(serverID, "config.json", conf, "file")

	embed = createCSGOMatchEmbed("Creating 1v1", "Unpacking info...", gameMap)
	s.ChannelMessageEditEmbed(m.ChannelID, status.ID, embed)
	csgo.UnzipPlugins(serverID)

	embed = createCSGOMatchEmbed("Creating 1v1", "Starting server...", gameMap)
	s.ChannelMessageEditEmbed(m.ChannelID, status.ID, embed)
	csgo.StartServer(serverID)

	serverIP, err := waitForServerIp(serverID)

	embed = createCSGOMatchEmbed("Creating 1v1", "Configuring server...", gameMap)
	s.ChannelMessageEditEmbed(m.ChannelID, status.ID, embed)

	err = csgo.SendCommandToServer(serverID, "get5_loadmatch config.json")

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to config server")
	}

	embed = createCSGOMatchEmbed("Playing 1v1", "`" + serverIP +"; password nadestack`", gameMap)
	s.ChannelMessageEditEmbed(m.ChannelID, status.ID, embed)




}

func waitForServerIp(id string) (string, error) {
	ch := make(chan string, 1)
	defer close(ch)

	go func() {
		for {
			if status, ip  := csgo.GetServerStatus(id); !status {
				ch<- ip
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	timer := time.NewTimer(2 * time.Minute)
	defer timer.Stop()

	select {
	case ip := <- ch:
		return ip, nil
	case <-timer.C:
		return "", errors.New("timed out waiting for serve to load")
	}

}
