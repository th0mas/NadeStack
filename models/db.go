package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Init postgres
	"github.com/th0mas/NadeStack/config"
)

// Models contains all the database models and helper methods
type Models struct {
	db *gorm.DB // Private database access hands off
}

// Init initializes the database according to the config file
func Init(c *config.Config) *Models {
	fmt.Println(c.DBUrl)
	db, err := gorm.Open("postgres", c.DBUrl+"?sslmode=disable")
	if err != nil {
		panic(err)
	}

	d := &Models{db}
	db.AutoMigrate(&User{}, &DeepLink{}, &Match{}, &Game{}) // what could go wrong
	return d
}

// Close closes the database connection when no longer needed
func (m *Models) Close() {
	_ = m.db.Close()
}
