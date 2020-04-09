package config

import (
	"github.com/spf13/viper"
)

// Config holds the configuration data for the bot
type Config struct {
	DiscordToken string `mapstructure:"discord_token"`   // DiscordToken is the Bot token used to authenticate with discord
	AzureToken   string `mapstructure:"azure_token"` // AzureToken is used to authenticate with azure to provision game servers
	Domain       string `mapstructure:"domain"`      // Domain is the domain and port the app is hosted on
	DBUrl        string `mapstructure:"database_url"`      // The URI for the postgres database
	DBType       string `mapstructure:"db_type"`     // The type of database either postgres or sqlite3
	Debug        bool   `mapstructure:"debug"`
}

// LoadConfig loads a yaml config from a given location
func LoadConfig(uri string) (c Config) {
	viper.SetDefault("debug", false)
	viper.SetConfigName("config.yml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	_ = viper.ReadInConfig()
	viper.BindEnv("discord_token", "DISCORD_TOKEN")
	viper.BindEnv("domain")
	viper.BindEnv("database_url", "DATABASE_URL")
	viper.BindEnv("debug")
	_ = viper.Unmarshal(&c)


	//err = yaml.UnmarshalStrict(dat, &c)
	//if err != nil {
	//	panic(err)
	//}

	return
}
