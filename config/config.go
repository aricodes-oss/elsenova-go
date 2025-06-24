package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token     string
	GuildID   string
	DailySeed struct {
		Channel  string
		Schedule string
	}

	DatabaseURL string

	BaseVoreCount     int
	SandwichFrequency float64
	VoreChannelId     string

	DebugLogging bool

	Cooldowns struct {
		Default int
		Vore    int
	}

	Web struct {
		Port     int
		Frontend string
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
