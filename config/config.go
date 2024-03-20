package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token   string
	GuildID string

	DatabaseURL   string
	BaseVoreCount int

	DebugLogging bool

	Cooldowns struct {
		Default int
		Vore    int
	}
}

func Load() *Config {
	ret := &Config{}
	viper.Unmarshal(ret)
	return ret
}
