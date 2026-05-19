package models

import "gorm.io/gorm"

type CustomCommand struct {
	gorm.Model

	Name        string `gorm:"index"`
	Description string
	Response    string
	CreatedBy   string
}

func init() {
	AllModels = append(AllModels, CustomCommand{})
}
