package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/th0mas/NadeStack/config"
)

type Models struct {
	db *gorm.DB // Private database access hands off
}

// Init initializes the database according to the config file
func Init(c *config.Config) *Models {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=tomh dbname=nadestack sslmode=disable")
	if err != nil {
		panic(err)
	}

	d := &Models{db}
	db.AutoMigrate(&User{}, &DeepLink{}) // what could go wrong
	return d
}

// Close closes the database connection when no longer needed
func (m *Models) Close() {
	_ = m.db.Close()
}
