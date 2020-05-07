package models

import (
	"github.com/jinzhu/gorm"
	"time"

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
	SkipVeto       bool           `json:"skip_veto"`
	SideType       string         `json:"side_type"`
	PlayersPerTeam int            `json:"players_per_team"`
	CVars          pq.StringArray `json:"cvars" gorm:"type:varchar(256)[]"`
	TeamOne        Team           `json:"team1"`
	TeamOneID      uint
	TeamTwo        Team `json:"team2"`
	TeamTwoID      uint
}

// Team is a team in a CsgoMatch. Needs a name and a list of players.
// Other fields that can be added in future/if needed are: tag, flag, logo
type Team struct {
	gorm.Model
	Name           string   `json:"name"`
	Players        []*User  `gorm:"many2many:user_teams;" json:"-"`
	PlayersSteamID []string `gorm:"-" json:"players, omitempty"`
	Score          *int
}

// Create5v5 creaes a 5v5 match config with specified maps and teams
func (m *Models) Create5v5(gameMap string, team1, team2 Team) *Match {
	match := &Match{
		ID:             createUniqueCode(),
		NumMaps:        1,
		MapList:        []string{gameMap},
		SkipVeto:       true,
		SideType:       "always_knife",
		PlayersPerTeam: 5,
		CVars:          []string{"get5_print_damage 1"},
		TeamOne:        team1,
		TeamTwo:        team2,
	}

	m.db.Create(match)
	return match
}

func (m *Models) Create1v1(gameMap string, team1, team2 Team) *Match {
	match := &Match{
		ID:             createUniqueCode(),
		NumMaps:        1,
		MapList:        []string{gameMap},
		SkipVeto:       true,
		SideType:       "always_knife",
		PlayersPerTeam: 1,
		CVars:          []string{"get5_print_damage 1"},
		TeamOne:        team1,
		TeamOneID:      team1.ID,
		TeamTwo:        team2,
		TeamTwoID:      team2.ID,
	}

	m.db.Create(match)
	return match
}

func (m *Models) GetMatchByID(id string) (*Match, error) {
	match := &Match{}
	err := m.db.Set("gorm:auto_preload", true).Where(&Match{ID: id}).First(match).Error
	//m.db.Model(match).Related(&match.TeamOne, "TeamOne")
	//m.db.Model(match).Related(&match.TeamTwo, "TeamTwo")

	return match, err
}

// CreateTeam initializes a csgo team
func (m *Models) CreateTeam(users []*User) *Team {
	t := &Team{
		Name:    users[0].DiscordNickname + "'s Team",
		Players: users,
	}

	return t
}

// GenerateTeamIDS geneates a list of Steam IDS for members of the team
func (m *Match) GenerateTeamIDS() {
	m.TeamOne.PlayersSteamID = make([]string, 0)
	m.TeamTwo.PlayersSteamID = make([]string, 0)

	for _, u := range m.TeamOne.Players {
		m.TeamOne.PlayersSteamID = append(m.TeamOne.PlayersSteamID, *u.SteamID)
	}
	for _, u := range m.TeamTwo.Players {
		m.TeamTwo.PlayersSteamID = append(m.TeamTwo.PlayersSteamID, *u.SteamID)
	}
}
