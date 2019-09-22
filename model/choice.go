package model

// Choice is the main subelement of the page
type Choice struct {
	ID            uint   `json:"id" 			gorm:"primary_key"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	ParentStoryID uint   `json:"parentStoryID" 	gorm:"column:parent_story_id"`
	NextStoryID   uint   `json:"nextStoryID" 	gorm:"column:next_story_id"`
	NextStory     Story  `json:"choices" 		gorm:"foreignkey:ID"`
}

func (choice Choice) ToDto() (ret struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	ParentStoryID uint   `json:"parentStoryID"`
	NextStoryID   uint   `json:"nextStoryID"`
}) {
	ret.ID = choice.ID
	ret.Title = choice.Title
	ret.Text = choice.Text
	ret.ParentStoryID = choice.ParentStoryID
	ret.NextStoryID = choice.NextStoryID

	return
}
