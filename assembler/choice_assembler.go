package assembler

import (
	model "../model"
)

// BuildChoiceDto converts an entity to a DTO
func BuildChoiceDto(entity model.Choice) model.ChoiceDto {
	var dto model.ChoiceDto

	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Text = entity.Text
	dto.ParentStoryRefer = entity.ParentStoryRefer
	dto.NextStoryRefer = entity.NextStoryRefer

	return dto
}

// BuildOrphanChoiceDto converts an entity to a DTO
func BuildOrphanChoiceDto(entity model.Choice) model.OrphanChoiceDto {
	var dto model.OrphanChoiceDto

	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Text = entity.Text
	dto.NextStoryRefer = entity.NextStoryRefer

	return dto
}

// BuildChoiceDto applies BuildChoiceDto on each element
func BuildChoicesDto(entities []model.Choice) []model.ChoiceDto {
	dtos := make([]model.ChoiceDto, len(entities))
	for i, entity := range entities {
		dtos[i] = BuildChoiceDto(entity)
	}
	return dtos
}

// BuildChoice converts a DTO to an entity
func BuildChoice(dto model.ChoiceDto) model.Choice {
	var entity model.Choice

	entity.Title = dto.Title
	entity.Text = dto.Text
	entity.ParentStoryRefer = dto.ParentStoryRefer
	entity.NextStoryRefer = dto.NextStoryRefer

	return entity
}
