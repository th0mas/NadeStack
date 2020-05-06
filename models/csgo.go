package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Match is a struct reflecting the config format for get_5 plugin, and should be encoded via JSON
// Possibly could persist some of these fields to database in future, so should mke sure ID fields etc are unique.
//
// https://github.com/splewis/get5#match-schema
type Match struct {
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      *time.Time     `json:"-"`
	ID             string         `json:"matchid" gorm:"PRIMARY_KEY"`
	NumMaps        int            `json:"num_maps"`
	MapList        pq.StringArray `json:"map_list" gorm:"type:varchar(50)[]"`
	SkipVeto       bool           `json:"skip_veto" gorm:"-"`
	SideType       string         `json:"side_type"`
	PlayersPerTeam int            `json:"players_per_team"`
	CVars          []string       `json:"cvars" gorm:"-"`
	TeamOne        Team           `json:"team1" gorm:"embedded;embedded_prefix:team_1_"`
	TeamTwo        Team           `json:"team2" gorm:"embedded;embedded_prefix:team_2_"`
	TeamOneScore   *int           `json:"-"`
	TeamTwoScore   *int           `json:"-"`
}

// Team is a team in a CsgoMatch. Needs a name and a list of players.
// Other fields that can be added in future/if needed are: tag, flag, logo
type Team struct {
	Name      string   `json:"name"`
	Players   []User   `gorm:"many2many:user_teams;"`
	PlayersID []string `gorm:"-" json:"players"`
}

type status int

const (
	notStarted status = iota
	serverUp
	configDeployed
	gameReady
	inProgress
	gameEnd
	gameOver
)

// Game is the lifecycle object for one csgo game
type Game struct {
	gorm.Model
	Match    Match
	MatchID  string
	Status   status
	ServerID *string
	ServerIP *string
}

// MakeGame creates a game instance given a match
func MakeGame(m *Match) *Game {
	return &Game{
		MatchID: m.ID,
		Status:  notStarted,
	}
}
