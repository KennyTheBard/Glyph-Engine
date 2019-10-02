package data

// ChoiceModel is the main subelement of the page
type ChoiceModel struct {
	BaseEntity
	Name               string `json:"name"`
	Text               string `json:"text"`
	ParentStoryID      uint   `json:"parentStoryID" 	gorm:"column:parent_story_id"`
	NextStoryScript    string `							gorm:"column:next_story_script"`
	DefaultNextStoryID uint   `json:"nextStoryID" 		gorm:"column:default_next_story_id"`
}

func (choice ChoiceModel) ToDto() (ret struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}) {
	ret.ID = choice.ID
	ret.Name = choice.Name
	ret.Text = choice.Text

	return
}

func (choice *ChoiceModel) GetNextStory() StoryModel {
	var story StoryModel
	DB.First(&story, choice.DefaultNextStoryID)
	return story
}

func (choice *ChoiceModel) GetItemStacks() []ItemStack {
	var stacks []ItemStack
	DB.Where("owner_id = ? and type = ?", choice.ID, OWNER_CHOICE).Find(&stacks)
	return stacks
}
