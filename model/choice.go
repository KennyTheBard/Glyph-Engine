package model

// ChoiceModel is the main subelement of the page
type ChoiceModel struct {
	ID            uint              `json:"id" 				gorm:"primary_key"`
	Name          string            `json:"name"`
	Text          string            `json:"text"`
	ParentStoryID uint              `json:"parentStoryID" 	gorm:"column:parent_story_id"`
	NextStoryID   uint              `json:"nextStoryID" 	gorm:"column:next_story_id"`
	NextStory     StoryModel        `json:"choices" 		gorm:"foreignkey:ID"`
	Costs         []ItemCost        `json:"costs" 			gorm:"one2many:ChoiceID"`
	Rewards       []ItemReward      `json:"rewards" 		gorm:"one2many:ChoiceID"`
	Requirements  []ItemRequirement `json:"requirements" 	gorm:"one2many:ChoiceID"`
}

func (choice ChoiceModel) ToDto() (ret struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Text          string `json:"text"`
	ParentStoryID uint   `json:"parentStoryID"`
	NextStoryID   uint   `json:"nextStoryID"`
}) {
	ret.ID = choice.ID
	ret.Name = choice.Name
	ret.Text = choice.Text
	ret.ParentStoryID = choice.ParentStoryID
	ret.NextStoryID = choice.NextStoryID

	return
}
