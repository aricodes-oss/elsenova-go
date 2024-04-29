package models

import (
	"fmt"
	"slices"
	"strings"

	"gorm.io/gorm"
)

type Sandwich struct {
	gorm.Model

	Name       string
	IsSandwich bool
}

// Classification returns either "sandwich" or "dumpling"
func (s *Sandwich) Classification() string {
	if s.IsSandwich {
		return "sandwich"
	}

	return "dumpling"
}

// Description returns a string of form "A(n) [object] is a [sandwich|dumpling]"
func (s *Sandwich) Description() string {
	article := "A"
	if vowels := []string{"a", "e", "i", "o", "u"}; slices.Contains(vowels, strings.ToLower(s.Name)[:1]) {
		article = "An"
	}

	return fmt.Sprintf("%s %s is a %s", article, s.Name, s.Classification())
}

func init() {
	AllModels = append(AllModels, Sandwich{})
}
