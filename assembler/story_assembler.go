package assembler

import (
	model "../model"
)

// StoryToDto converts an entity to a DTO
func StoryToDto(entity model.Story) model.StoryDto {
	var dto model.StoryDto

	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Text = entity.Text

	return dto
}
