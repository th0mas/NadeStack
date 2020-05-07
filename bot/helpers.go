package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

func findUserVoiceChannel(s *discordgo.Session, guildID, userID string) (string, error) {
	g, _ := s.State.Guild(guildID)

	for _, vs := range g.VoiceStates {
		if vs.UserID == userID {
			return vs.ChannelID, nil
		}
	}

	return "", errors.New("could not find user in a channel")
}
