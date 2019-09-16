package model

import (
	"github.com/jinzhu/gorm"
)

// Story is the main element of a page
type Story struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}

// StoryDto encapsulates fields that should be seen by others
type StoryDto struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Choices []uint `json:"choices"`
}
