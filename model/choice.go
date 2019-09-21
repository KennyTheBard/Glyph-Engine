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

func (choice Choice) ToDto() (ret struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	Text             string `json:"text"`
	ParentStoryRefer uint   `json:"parentStoryID"`
	NextStoryRefer   uint   `json:"nextStoryID"`
}) {
	ret.ID = choice.ID
	ret.Title = choice.Title
	ret.Text = choice.Text
	ret.ParentStoryRefer = choice.ParentStoryRefer
	ret.NextStoryRefer = choice.NextStoryRefer

	return
}
