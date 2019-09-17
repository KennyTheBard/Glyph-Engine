package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	asm "../assembler"
	data "../data"
	model "../model"
	service "../service"
)

// CreateChoice creates a choice
func CreateChoice(context *gin.Context) {
	var dto model.ChoiceDto
	context.BindJSON(&dto)

	choice, err := service.SaveChoice(asm.BuildChoice(dto))
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Choice created successfully!", "resourceId": choice.ID})
}

// GetAllChoices retrieves all choices
func GetAllChoices(context *gin.Context) {
	choices := service.FindAllChoices()
	dtos := asm.BuildChoicesDto(choices)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetChoice retrieves a choice
func GetChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	choice, err := service.FindChoiceById(uint(id))
	if err != nil {
		panic(err)
	}

	dto := asm.BuildChoiceDto(choice)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateChoice updates a choice
func UpdateChoice(context *gin.Context) {
	var dto model.ChoiceDto
	context.BindJSON(&dto)

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	choice, err := service.FindChoiceById(uint(id))
	if err != nil {
		panic(err)
	}

	data.DB.Model(&choice).Update("title", dto.Title)
	data.DB.Model(&choice).Update("text", dto.Text)
	data.DB.Model(&choice).Update("parent_story", dto.ParentStory)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice updated successfully!"})
}

// DeleteChoice removes a choice
func DeleteChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	err = service.DeleteChoiceById(uint(id))
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice deleted successfully!"})
}
