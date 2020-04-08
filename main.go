package main

import (
	"flag"
	"github.com/th0mas/NadeStack/bot"
	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"
	"github.com/th0mas/NadeStack/web"
	"os"
	"os/signal"
	"syscall"
)

// Command line params
var (
	ConfigPath string
)

type Service interface {
	Run(c *config.Config, db *models.DB)
	Close()
}

func init() {
	flag.StringVar(&ConfigPath, "c", "config.yml", "Filepath to the config file")

	flag.Parse()
}

func main() {
	c := config.LoadConfig(ConfigPath)
	db := models.Init(&c)

	// Run component services
	// writing like this allows expansion in future.
	services := []Service{
		&web.Web{},
		&bot.Bot{},
	}

	for _, s := range services {
		s.Run(&c, db)
	}

	// Safely close connections on close
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	for _, s := range services {
		s.Close()
	}
}
