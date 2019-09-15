package model

import (
	"github.com/jinzhu/gorm"
)

// Choice is the main subelement of the page
type Choice struct {
	gorm.Model
	Title   string  `json:"title"`
	Text    string  `json:"text"`
	Stories []Story `json:"stories" gorm:"many2many:stories;"`
}

// ChoiceDto encapsulates fields that should be seen by others
type ChoiceDto struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Stories []uint `json:"stories"`
}
