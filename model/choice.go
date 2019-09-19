package model

// Choice is the main subelement of the page
type Choice struct {
	ID               uint `gorm:"primary_key"`
	Title            string
	Text             string
	ParentStoryRefer uint  `gorm:"column:parent_story_refer"`
	NextStoryRefer   uint  `gorm:"column:next_story_refer"`
	NextStory        Story `gorm:"foreignkey:ID"`
}

// ChoiceDto encapsulates fields that should be seen by others + the parent
type ChoiceDto struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	Text             string `json:"text"`
	ParentStoryRefer uint   `json:"parentStoryID"`
	NextStoryRefer   uint   `json:"nextStoryID"`
}

// OrphanChoiceDto encapsulates fields that should be seen by others
type OrphanChoiceDto struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Text           string `json:"text"`
	NextStoryRefer uint   `json:"nextStoryID"`
}
