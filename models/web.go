package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Action is a type defing a deeplink action
type Action int

const (
	// LinkSteamID asks to link a steam ID to a user account
	LinkSteamID Action = iota
)

const (
	notFoundError       = "models: user not found"
	discordAvatarCdnUrl = "https://cdn.discordapp.com/avatars/"
)

type User struct {
	gorm.Model
	DiscordID            string  `gorm:"unique"`
	SteamID              *string `gorm:"unique"`
	SteamID64            *uint64 `gorm:"unique"`
	DiscordNickname      string
	DiscordProfilePicURL string
	Games                []*Team `gorm:"many2many:user_teams;"`
}

type DeepLink struct {
	gorm.Model
	ShortURL   string `gorm:"unique"`
	UserID     uint
	User       User
	LinkAction Action
	Payload    interface{} `gorm:"-"`
}

func (m *Models) GetUserByDiscordID(discordID string) (u *User, err error) {
	u = &User{}
	err = m.db.Where(&User{DiscordID: discordID}).First(u).Error
	return u, err
}

func (m *Models) CreateUserFromDiscord(discordID, discordNickname, avatarHash string) *User {
	u := &User{
		DiscordID:            discordID,
		DiscordNickname:      discordNickname,
		DiscordProfilePicURL: discordAvatarCdnUrl + fmt.Sprintf("%s/%s.png", discordID, avatarHash),
	}

	m.db.Create(u)
	return u
}

func (m *Models) AddSteamIDToUser(u *User, steamID string, steamID64 uint64) {
	m.db.Model(&u).Updates(User{SteamID: &steamID, SteamID64: &steamID64})

	return
}

func (m *Models) UpdateDiscordInfo(u *User, nickname, avatar string) {
	m.db.Model(&u).Updates(User{DiscordNickname: nickname, DiscordProfilePicURL: discordAvatarCdnUrl + fmt.Sprintf("%s/%s.png", u.DiscordID, avatar)})
}
func (m *Models) CreateDeepLink(action Action, u *User) *DeepLink {
	code := createUniqueCode()
	d := &DeepLink{
		ShortURL:   code,
		UserID:     u.ID,
		LinkAction: action,
	}
	m.db.Create(d)
	return d
}

// TODO: With no rune this exposes data
func (m *Models) GetDeepLinkData(rune string) (d *DeepLink, err error) {
	d = &DeepLink{}
	err = m.db.Where(&DeepLink{ShortURL: rune}).First(d).Error
	m.db.Model(d).Related(&d.User)
	return d, err
}

// IsRecordNotFound is a wrapper around the built in function
func (m *Models) IsRecordNotFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

// https://stackoverflow.com/a/39482484
func createUniqueCode() string {
	c := 8
	b := make([]byte, c)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(b)
}
