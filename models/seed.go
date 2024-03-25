package models

import (
	"gorm.io/gorm"
)

type Seed struct {
	gorm.Model

	Value string `json:"value" gorm:"unique"`
}
