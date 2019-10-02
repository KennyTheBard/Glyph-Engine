package data

// Story is the main element of a page
type StoryModel struct {
	BaseEntity
	Name string `json:"name"`
	Text string `json:"text"`
}

func (story StoryModel) ToDto() (ret struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}) {
	ret.ID = story.ID
	ret.Name = story.Name
	ret.Text = story.Text

	return
}

func (story *StoryModel) GetChoices() []ChoiceModel {
	var choices []ChoiceModel
	DB.Where("parent_story_id = ?", story.ID).Find(&choices)
	return choices
}
