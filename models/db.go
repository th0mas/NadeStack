package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/th0mas/NadeStack/config"
)

type Models struct {
	db *gorm.DB // Private database access hands off
}

// Init initializes the database according to the config file
func Init(c *config.Config) *Models {
	//u, err := url.Parse(c.DBUrl)
	//if err != nil {
	//	panic("could not parse db url")
	//}
	//var connStr string
	//fmt.Printf("%s", u.Path)
	//if v, p := u.User.Password() ; p {
	//	connStr = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
	//		u.User.Username(), v, u.Host, u.Port(), u.Path)
	//} else {
	//	connStr = fmt.Sprintf("user=%s host=%s port=%s dbname=%s",
	//		u.User.Username(), u.Host, u.Port(), u.Path)
	//}
	fmt.Println("Starting Database")
	db, err := gorm.Open("postgres", c.DBUrl + "?sslmode=disable")
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
