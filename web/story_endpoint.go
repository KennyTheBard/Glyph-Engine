package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	asm "../assembler"
	data "../data"
	model "../model"
)

// CreateStory creates a story
func CreateStory(context *gin.Context) {
	var dto model.StoryDto
	context.BindJSON(&dto)
	story := model.Story{Title: dto.Title, Text: dto.Text}
	data.DB.Save(&story)

	for _, choiceID := range dto.Choices {
		var choice model.Choice
		data.DB.First(&choice, choiceID).Update("parent_id", story.ID)
		data.DB.Save(&choice) // THIS WAS THE THINGY
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Story created successfully!", "resourceId": story.ID})
}

// GetAllStories retrieves all stories
func GetAllStories(context *gin.Context) {
	var stories []model.Story
	data.DB.Find(&stories)

	var dtos []model.StoryDto
	for _, item := range stories {
		dtos = append(dtos, asm.StoryToDto(item))
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetStory retrieves a story
func GetStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")
	data.DB.First(&story, id)

	dto := asm.StoryToDto(story)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var updatedStory model.StoryDto
	context.BindJSON(&updatedStory)

	var story model.Story
	id := context.Param("id")
	data.DB.First(&story, id)

	if story.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No story found!"})
		return
	}

	var choices []model.Choice
	for _, choiceID := range updatedStory.Choices {
		var choice model.Choice
		data.DB.First(&choice, choiceID)
		choices = append(choices, choice)
	}

	fmt.Println(choices)

	data.DB.Model(&story).Update("title", updatedStory.Title)
	data.DB.Model(&story).Update("text", updatedStory.Text)
	data.DB.Model(&story).Update("choices", choices)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story updated successfully!"})
}

// DeleteStory removes a story
func DeleteStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")

	data.DB.First(&story, id)

	if story.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No story found!"})
		return
	}

	data.DB.Delete(&story)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story deleted successfully!"})
}
