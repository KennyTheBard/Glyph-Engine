package service

import (
	"errors"

	data "../data"
	model "../model"
)

func SaveStory(story model.Story) (model.Story, error) {
	data.DB.Save(&story)
	return story, nil
}

func FindStoryById(id uint) (model.Story, error) {
	var story model.Story
	data.DB.First(&story, id)
	if story.ID != id {
		return story, errors.New("No story found with the given ID")
	}
	return story, nil
}

func FindAllStories() []model.Story {
	var stories []model.Story
	data.DB.Find(&stories)
	return stories
}

func DeleteStoryById(id uint) error {
	var story model.Story
	data.DB.First(&story, id)
	if story.ID != id {
		return errors.New("No story found with the given ID")
	}
	data.DB.Delete(&story)
	return nil
}
