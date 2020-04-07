package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration data for the bot
type Config struct {
	DiscordToken string `yaml:"bot_token"` // DiscordToken is the Bot token used to authenticate with discord
	AzureToken   string `yaml:"azure_token"` // AzureToken is used to authenticate with azure to provision game servers
	Domain		 string `yaml:"domain"` // Domain is the domain and port the app is hosted on
}

// LoadConfig loads a yaml config from a given location
func LoadConfig(uri string) (c Config) {
	dat, err := ioutil.ReadFile(uri)
	if err != nil {
		panic(err)
	}

	err = yaml.UnmarshalStrict(dat, &c)
	if err != nil {
		panic(err)
	}

	return
}
