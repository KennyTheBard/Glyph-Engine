package model

// ChoiceModel is the main subelement of the page
type ChoiceModel struct {
	ID            uint        `json:"id" 				gorm:"primary_key"`
	Name          string      `json:"name"`
	Text          string      `json:"text"`
	EnergyCost    uint        `json:energyCost`
	ParentStoryID uint        `json:"parentStoryID" 	gorm:"column:parent_story_id"`
	NextStoryID   uint        `json:"nextStoryID" 		gorm:"column:next_story_id"`
	NextStory     StoryModel  `json:"nextStory" 		gorm:"foreignkey:ID"`
	Stacks        []ItemStack `json:"stacks" 			gorm:"one2many:ChoiceID"`
}

func (choice ChoiceModel) ToDto() (ret struct {
	ID            uint        `json:"id"`
	Name          string      `json:"name"`
	Text          string      `json:"text"`
	EnergyCost    uint        `json:energyCost`
	ParentStoryID uint        `json:"parentStoryID"`
	NextStoryID   uint        `json:"nextStoryID"`
	Stacks        []ItemStack `json:"stacks"`
}) {
	ret.ID = choice.ID
	ret.Name = choice.Name
	ret.Text = choice.Text
	ret.EnergyCost = choice.EnergyCost
	ret.ParentStoryID = choice.ParentStoryID
	ret.NextStoryID = choice.NextStoryID
	ret.Stacks = choice.Stacks

	return
}
