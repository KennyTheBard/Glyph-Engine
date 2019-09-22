package model

// Story is the main element of a page
type StoryModel struct {
	ID      uint          `json:"id" 		gorm:"primary_key"`
	Name    string        `json:"name"`
	Text    string        `json:"text"`
	Choices []ChoiceModel `json:"choices" 	gorm:"one2many:parentStoryID"`
}

func (story StoryModel) ToDto() (ret struct {
	ID      uint          `json:"id"`
	Name    string        `json:"name"`
	Text    string        `json:"text"`
	Choices []ChoiceModel `json:"choices"`
}) {
	ret.ID = story.ID
	ret.Name = story.Name
	ret.Text = story.Text
	ret.Choices = story.Choices

	return
}
