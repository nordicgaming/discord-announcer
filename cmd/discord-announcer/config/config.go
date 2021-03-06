package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config for EventStore and Discord
type Config struct {
	Discord struct {
		BotToken  string
		UserToken string
		Username  string
		Password  string
	}
}

// ReadConfig from config files or environment
func ReadConfig(cfg *Config) error {
	v := viper.New()
	v.SetEnvPrefix("Announcer")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetDefault("discord.bottoken", "")
	v.SetDefault("discord.usertoken", "")
	v.SetDefault("discord.username", "")
	v.SetDefault("discord.password", "")

	return v.Unmarshal(cfg)
}
