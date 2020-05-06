package models

import "github.com/jinzhu/gorm"

// Game is the lifecycle object for one csgo game
type Game struct {
	gorm.Model
	Match    Match
	MatchID  string
	Status   status
	ServerID *string
	ServerIP *string
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

// MakeGame creates a game instance given a match
func (m *Models) MakeGame(match *Match) *Game {
	g := &Game{
		MatchID: match.ID,
		Status:  notStarted,
	}

	m.db.Create(g)
	return g

}

// GetGame gets a game by its integer ID
func (m *Models) GetGame(id int) (*Game, error) {
	g := &Game{}
	err := m.db.First(g, id).Error

	return g, err
}
