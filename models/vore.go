package models

import (
	"gorm.io/gorm"
)

type Vore struct {
	gorm.Model

	UserID string `json:"userID"`
}
