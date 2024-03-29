package models

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type CachedUser struct {
	gorm.Model
	discordgo.User

	ID string `json:"id" gorm:"primarykey"`
}

func init() {
	AllModels = append(AllModels, CachedUser{})
}
