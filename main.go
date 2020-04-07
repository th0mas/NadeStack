package main

import (
	"flag"
	"github.com/th0mas/NadeStack/bot"
	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/web"
	"os"
	"os/signal"
	"syscall"
)

// Command line params
var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "c", "config.yml", "Filepath to the config file")

	flag.Parse()
}

func main() {
	c := config.LoadConfig(ConfigPath)

	// Run component services
	discordBot := bot.Run(&c)
	web.Run(&c)

	// Safely close connections on close
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// safely close discod
	_ = discordBot.Close()
}
