package data

import (
	"errors"

	model "../model"
)

func SaveStory(story model.StoryModel) (model.StoryModel, error) {
	DB.Save(&story)
	return story, nil
}

func FindStoryById(id uint) (model.StoryModel, error) {
	var story model.StoryModel
	if id == 0 {
		return story, errors.New("ID's must be positive numbers")
	}

	DB.First(&story, id)
	if story.ID != id {
		return story, errors.New("No story found with the given ID")
	}

	var choices []model.ChoiceModel
	DB.Model(&story).Related(&choices, "ParentStoryID")
	story.Choices = choices
	return story, nil
}

func FindAllStories() []model.StoryModel {
	var stories []model.StoryModel
	DB.Find(&stories)
	return stories
}

func UpdateStoryField(story model.StoryModel, fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(&story).Update(name, value)
	}

	return nil
}

func DeleteStory(story model.StoryModel) error {
	DB.Delete(&story)
	return nil
}
