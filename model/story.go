package model

import (
	"github.com/jinzhu/gorm"
)

type Story struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}

type StoryDto struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
