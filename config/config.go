package config

import (
	"github.com/spf13/viper"
)

// Config holds the configuration data for the bot
type Config struct {
	DiscordToken string `mapstructure:"discord_token"` // DiscordToken is the Bot token used to authenticate with discord
	DatHostUser  string `mapstructure:"dat_host_user"` // Username for the dathost API
	DatHostPass  string `mapstructure:"dat_host_pass"` // Password for the dathost api
	GSLT         string `mapstructure:"gslt"`
	Domain       string `mapstructure:"domain"`       // Domain is the domain and port the app is hosted on
	DBUrl        string `mapstructure:"database_url"` // The URI for the postgres database
	DBType       string `mapstructure:"db_type"`      // The type of database either postgres or sqlite3
	Debug        bool   `mapstructure:"debug"`
	CmdPrefix    string `mapstructure:"cmd_prefix"`
}

// LoadConfig loads a yaml config from a given location
func LoadConfig(uri string) (c Config) {
	viper.SetDefault("debug", false)
	viper.SetDefault("cmd_prefix", "/")
	viper.SetConfigName("config.yml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	_ = viper.ReadInConfig()
	viper.BindEnv("discord_token", "DISCORD_TOKEN")
	viper.BindEnv("domain")
	viper.BindEnv("database_url", "DATABASE_URL")
	viper.BindEnv("gslt", "GSLT")
	viper.BindEnv("dat_host_user", "DAT_HOST_USER")
	viper.BindEnv("dat_host_pass", "DAT_HOST_PASS")
	viper.BindEnv("debug")
	_ = viper.Unmarshal(&c)

	//err = yaml.UnmarshalStrict(dat, &c)
	//if err != nil {
	//	panic(err)
	//}

	return
}
