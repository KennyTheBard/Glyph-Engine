package data

import "errors"

// Story is the main element of a page
type StoryModel struct {
	ID   uint   `json:"id" gorm:"primary_key"`
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

// CRUD methods

func (story *StoryModel) Save() error {
	DB.Save(story)
	return nil
}

func (story *StoryModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(story, id)
	if story.ID != id {
		return errors.New("No story found with the given ID")
	}

	return nil
}

func (story *StoryModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(story).Update(fieldName, fieldValue)

	return nil
}

func (story *StoryModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(story).Update(name, value)
	}

	return nil
}

func (story *StoryModel) Delete() error {
	DB.Delete(story)
	return nil
}
