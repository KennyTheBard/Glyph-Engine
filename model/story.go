package model

// Story is the main element of a page
type Story struct {
	ID      uint `gorm:"primary_key"`
	Title   string
	Text    string
	Choices []Choice `gorm:"one2many:ParentStoryRefer"`
}

// StoryDto encapsulates fields that should be seen by others
type StoryDto struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// StoryCompleteDto encapsulates fields that should be seen by others + the available choices
type StoryCompleteDto struct {
	ID      uint              `json:"id"`
	Title   string            `json:"title"`
	Text    string            `json:"text"`
	Choices []OrphanChoiceDto `json:"choices"`
}
