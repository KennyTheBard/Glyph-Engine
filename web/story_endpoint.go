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
	var story model.Story
	context.BindJSON(&story)
	data.DB.Save(&story)

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Story created successfully!", "resourceId": story.ID})
}

// GetAllStories retrieves all stories
func GetAllStories(context *gin.Context) {
	var stories []model.Story
	data.DB.Find(&stories)

	var dtos []model.StoryDto
	for _, item := range stories {
		fmt.Print(">>>>> ")
		fmt.Println(item)
		dtos = append(dtos, asm.StoryToDto(item))
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetStory retrieves a story
func GetStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")
	data.DB.First(&story, id)

	fmt.Print(">>>>> ")
	fmt.Println(story)

	dto := asm.StoryToDto(story)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var updatedStory model.StoryDto
	context.BindJSON(&updatedStory)

	fmt.Print(">>>>> ")
	fmt.Println(updatedStory)

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
		fmt.Print(">>>>> ")
		fmt.Println(choice)
		choices = append(choices, choice)
	}

	data.DB.Model(&story).Update("title", updatedStory.Title)
	data.DB.Model(&story).Update("text", updatedStory.Text)
	data.DB.Model(&story).Update("choices", choices)

	fmt.Print(">>>>> ")
	fmt.Println(story)

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
