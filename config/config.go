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

	Web struct {
		Port int
	}
}

var current *Config

func Load() *Config {
	if current != nil {
		return current
	}

	current = &Config{}
	viper.Unmarshal(current)
	return current
}