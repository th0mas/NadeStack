package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type Action int

const (
	LinkSteamID Action = iota
)

const (
	userNotFoundError = "models: user not found"
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

func (m *Models) GetUserByDiscordID(discordID string) (*User, error) {
	var err error
	u := new(User)
	if err = m.db.Where(&User{DiscordID: discordID}).First(u).Error; gorm.IsRecordNotFoundError(err) {
		err = errors.New(userNotFoundError)
	}
	return u, err
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

func (m *Models) getDeepLinkData(rune string) (d *DeepLink) {
	m.db.Where(&DeepLink{ShortURL: rune}).First(d)
	return
}

func (m *Models) CheckUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasSuffix(err.Error(), userNotFoundError)
}
