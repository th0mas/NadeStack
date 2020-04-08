package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/th0mas/NadeStack/config"
)

type DB struct {
	db *sql.DB // Private database access hands off
}

// Init initializes the database according to the config file
func Init(c *config.Config) *DB {
	db, err := sql.Open("postgres", c.DBUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	d := &DB{db}

	return d
}

// Close closes the database connection when no longer needed
func (d *DB) Close() {
	_ = d.db.Close()
}
