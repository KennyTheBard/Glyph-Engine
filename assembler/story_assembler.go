package assembler

import (
	model "../model"
)

// BuildStoryDto converts an entity to a DTO
func BuildStoryDto(entity model.Story) model.StoryDto {
	var dto model.StoryDto

	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Text = entity.Text

	return dto
}

// BuildCompleteStoryDto converts an entity to a DTO
func BuildCompleteStoryDto(entity model.Story, choices []model.Choice) model.StoryCompleteDto {
	var dto model.StoryCompleteDto

	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Text = entity.Text
	dto.Choices = make([]model.OrphanChoiceDto, len(choices))
	for i, choice := range choices {
		dto.Choices[i] = BuildOrphanChoiceDto(choice)
	}

	return dto
}

// BuildStoriesDto applies BuildStoryDto on each element
func BuildStoriesDto(entities []model.Story) []model.StoryDto {
	dtos := make([]model.StoryDto, len(entities))
	for i, entity := range entities {
		dtos[i] = BuildStoryDto(entity)
	}
	return dtos
}

// BuildStory converts a DTO to an entity
func BuildStory(dto model.StoryDto) model.Story {
	var entity model.Story

	entity.Title = dto.Title
	entity.Text = dto.Text

	return entity
}
