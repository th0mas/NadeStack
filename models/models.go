package models

import "github.com/jinzhu/gorm"

type Action int

const (
	LinkSteamID Action = iota
)

type User struct {
	gorm.Model
	DiscordID            string `gorm:"unique"`
	SteamID              string `gorm:"unique"`
	SteamID64            uint64 `gorm:"unique"`
	DiscordNickname      string
	DiscordProfilePicURL string
}

type DeepLink struct {
	gorm.Model
	ShortURL   string `gorm:"unique"`
	UserID     uint
	User       User
	LinkAction Action
}

func (m *Models) GetUserByDiscordID(discordID string) (u *User) {
	m.db.Where(&User{DiscordID: discordID}).First(u)
	return
}

func (m *Models) CreateUserFromDiscord(discordID, discordNickname, discordProfilePicURL string) *User {
	u := &User{
		DiscordID:            discordID,
		DiscordNickname:      discordNickname,
		DiscordProfilePicURL: discordProfilePicURL,
	}

	m.db.Create(u)
	return u
}

func (m *Models) AddSteamIDToUser(u *User, steamID string, steamID64 uint64) {
	m.db.Model(&u).Updates(User{SteamID: steamID, SteamID64: steamID64})

	return
}
