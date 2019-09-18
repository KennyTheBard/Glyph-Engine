package model

import (
	"github.com/jinzhu/gorm"
)

// Choice is the main subelement of the page
type Choice struct {
	gorm.Model
	Title       string `json:"title"`
	Text        string `json:"text"`
	ParentStory uint   `json:"parent_story"`
	NextStory   uint   `json:"next_story"`
}

// ChoiceDto encapsulates fields that should be seen by others + the parent
type ChoiceDto struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	ParentStory uint   `json:"parent_story"`
	NextStory   uint   `json:"next_story"`
}

// OrphanChoiceDto encapsulates fields that should be seen by others
type OrphanChoiceDto struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	NextStory uint   `json:"next_story"`
}
