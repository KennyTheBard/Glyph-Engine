package model

import (
	"github.com/jinzhu/gorm"
)

// Item is the principal form of resource in the game
type Item struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}

// ItemDto encapsulates fields that should be seen by others
type ItemDto struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
