package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
)

// CreateChoice creates a choice
func CreateChoice(context *gin.Context) {
	var dto model.ChoiceDto
	context.BindJSON(&dto)

	choice := model.Choice{
		Title: dto.Title,
		Text:  dto.Text}
	data.DB.Save(&choice)

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Choice created successfully!", "resourceId": choice.ID})
}

// GetAllChoices retrieves all choices
func GetAllChoices(context *gin.Context) {
	var choices []model.Choice
	data.DB.Find(&choices)

	var dtos []model.ChoiceDto
	for _, item := range choices {
		dtos = append(dtos, model.ChoiceDto{ID: item.ID, Title: item.Title, Text: item.Text})
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetChoice retrieves a choice
func GetChoice(context *gin.Context) {
	var choice model.Choice
	id := context.Param("id")
	data.DB.First(&choice, id)

	dto := model.ChoiceDto{ID: choice.ID, Title: choice.Title, Text: choice.Text}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateChoice updates a choice
func UpdateChoice(context *gin.Context) {
	var updatedChoice model.Choice
	var choice model.Choice
	context.BindJSON(&updatedChoice)

	id := context.Param("id")
	data.DB.First(&choice, id)

	if choice.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No choice found!"})
		return
	}

	data.DB.Model(&choice).Update("title", updatedChoice.Title)
	data.DB.Model(&choice).Update("text", updatedChoice.Text)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice updated successfully!"})
}

// DeleteChoice removes a choice
func DeleteChoice(context *gin.Context) {
	var choice model.Choice
	id := context.Param("id")

	data.DB.First(&choice, id)

	if choice.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No choice found!"})
		return
	}

	data.DB.Delete(&choice)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice deleted successfully!"})
}
