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

	e := createLinkEmbed("Link Steam Account", "To be able to use NadeStack, you must first link your Steam account.", fmt.Sprintf("%s/%s", b.conf.Domain, dl.ShortURL))

	s.ChannelMessageSendEmbed(userChannel.ID, e)

	s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
}

func (b *Bot) profileCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u, err := b.models.GetUserByDiscordID(m.Author.ID)
	if err != nil || u.SteamID == nil {
		s.ChannelMessageSend(m.ChannelID, "no user infomation found for discord id: "+m.Author.ID)
		return
	}

	e := createUserInfoEmbed(u)
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, e)

	if err != nil {
		panic(err)
	}

}

func (b *Bot) updateCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	u, err := b.models.GetUserByDiscordID(m.Author.ID)
	if err != nil || u.SteamID == nil {
		s.ChannelMessageSend(m.ChannelID, "no user infomation found for discord id: "+m.Author.ID)
		return
	}

	b.models.UpdateDiscordInfo(u, m.Author.Username, m.Author.Avatar)
}

func (b *Bot) start(s *discordgo.Session, m *discordgo.MessageCreate) {
	// start the server - info
	e := createCSGOMatchEmbed("Creating CSGO server...", "This could take up to a minute to complete", "de_mirage")
	message, err := s.ChannelMessageSendEmbed(m.ChannelID, e)
	if err != nil {
		fmt.Println(message)
		panic(err)
	}

}
