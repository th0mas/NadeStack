package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/th0mas/NadeStack/models"
)

func (b *Bot) steamLinkCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u, err := b.models.GetUserByDiscordID(m.Author.ID)
	if b.models.IsRecordNotFound(err) {
		if b.conf.Debug {
			s.ChannelMessageSend(m.ChannelID, "`DEBUG: No user with Discord ID "+m.Author.ID+" creating user")
		}
		u = b.models.CreateUserFromDiscord(m.Author.ID, m.Author.Username, m.Author.Avatar)
	} else {
		if b.conf.Debug {
			s.ChannelMessageSend(m.ChannelID, "`DEBUG: User already exists, skipping create`")
		}
	}
	userChannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		panic(err)
	}

	dl := b.models.CreateDeepLink(models.LinkSteamID, u)
	if b.conf.Debug {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`DEBUG: Created Deep Link with vals %+v`", dl))
	}
	s.ChannelMessageSend(userChannel.ID, fmt.Sprintf("Click the link to link you accounts: %s/%s", b.conf.Domain, dl.ShortURL))

	s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
}

func (b *Bot) steamDebugCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u, err := b.models.GetUserByDiscordID(m.Author.ID)
	if err != nil || u.SteamID == nil {
		s.ChannelMessageSend(m.ChannelID, "no user infomation found for discord id: "+m.Author.ID)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Steam connection infomation: \n SteamID: `%s` \n SteamID64: '%d'",
		*u.SteamID, *u.SteamID64))
}
