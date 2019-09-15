package assembler

import (
	model "../model"
)

// ChoiceToDto converts an entity to a DTO
func ChoiceToDto(entity model.Choice) model.ChoiceDto {
	var dto model.ChoiceDto

	dto.ID = entity.ID
	dto.Title = entity.Title
	dto.Text = entity.Text

	return dto
}
