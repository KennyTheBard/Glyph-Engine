package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	asm "../assembler"
	data "../data"
	model "../model"
)

// CreateChoice creates a choice
func CreateChoice(context *gin.Context) {
	var choice model.Choice
	var err error
	dto := choice.ToDto()
	if err = context.BindJSON(&dto); err != nil {
		StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	choice, err = data.SaveChoice(asm.BuildChoice(dto))
	if err != nil {
		StatusResponse(context, http.StatusInternalServerError, "Failed to create new choice!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Choice created successfully!", "resourceId": choice.ID})
}

// GetAllChoices retrieves all choices
func GetAllChoices(context *gin.Context) {
	choices := data.FindAllChoices()
	dtos := asm.BuildChoicesDto(choices)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetChoice retrieves a choice
func GetChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	choice, err := service.FindChoiceById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	dto := asm.BuildChoiceDto(choice)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateChoice updates a choice
func UpdateChoice(context *gin.Context) {
	var dto model.ChoiceDto
	if err := context.BindJSON(&dto); err != nil {
		StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}
	choice, err := service.FindChoiceById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	data.DB.Model(&choice).Update("title", dto.Title)
	data.DB.Model(&choice).Update("text", dto.Text)
	data.DB.Model(&choice).Update("parent_story_refer", dto.ParentStoryRefer)
	data.DB.Model(&choice).Update("next_story_refer", dto.NextStoryRefer)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice updated successfully!"})
}

// DeleteChoice removes a choice
func DeleteChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}
	err = service.DeleteChoiceById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice deleted successfully!"})
}
