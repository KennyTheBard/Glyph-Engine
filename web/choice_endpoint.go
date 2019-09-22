package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
)

// CreateChoice creates a choice
func CreateChoice(context *gin.Context) {
	var choice model.Choice
	if err := context.BindJSON(&choice); err != nil {
		StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	choice, err := data.SaveChoice(choice)
	if err != nil {
		StatusResponse(context, http.StatusInternalServerError, "Failed to create new choice!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Choice created successfully!", "resourceId": choice.ID})
}

// GetAllChoices retrieves all choices
func GetAllChoices(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": data.FindAllChoices()})
}

// GetChoice retrieves a choice
func GetChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	choice, err := data.FindChoiceById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": choice.ToDto()})
}

// UpdateChoice updates a choice
func UpdateChoice(context *gin.Context) {
	var updateChoice model.Choice
	if err := context.BindJSON(&updateChoice); err != nil {
		StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}
	choice, err := data.FindChoiceById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	data.DB.Model(&choice).Update("title", updateChoice.Title)
	data.DB.Model(&choice).Update("text", updateChoice.Text)
	data.DB.Model(&choice).Update("parent_story_id", updateChoice.ParentStoryID)
	data.DB.Model(&choice).Update("next_story_id", updateChoice.NextStoryID)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice updated successfully!"})
}

// DeleteChoice removes a choice
func DeleteChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}
	err = data.DeleteChoiceById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice deleted successfully!"})
}
