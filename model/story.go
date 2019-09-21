package model

// Story is the main element of a page
type Story struct {
	ID      uint     `gorm:"primary_key" json:"id"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Choices []Choice `gorm:"one2many:ParentStoryRefer"`
}

func (story Story) ToDto() (ret struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Choices []Choice `json:"choices"`
}) {
	ret.ID = story.ID
	ret.Title = story.Title
	ret.Text = story.Text
	ret.Choices = story.Choices

	return
}
