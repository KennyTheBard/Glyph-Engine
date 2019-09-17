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

// CreateStory creates a story
func CreateStory(context *gin.Context) {
	var dto model.StoryDto
	context.BindJSON(&dto)

	story, err := service.SaveStory(asm.BuildStory(dto))
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Story created successfully!", "resourceId": story.ID})
}

// GetAllStories retrieves all stories
func GetAllStories(context *gin.Context) {
	stories := service.FindAllStories()
	dtos := asm.BuildStoriesDto(stories)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetStory retrieves a story
func GetStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	story, err := service.FindStoryById(uint(id))
	if err != nil {
		panic(err)
	}

	dto := asm.BuildStoryDto(story)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var dto model.StoryDto
	context.BindJSON(&dto)

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	story, err := service.FindStoryById(uint(id))
	if err != nil {
		panic(err)
	}

	data.DB.Model(&story).Update("title", dto.Title)
	data.DB.Model(&story).Update("text", dto.Text)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story updated successfully!"})
}

// DeleteStory removes a story
func DeleteStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	err = service.DeleteStoryById(uint(id))
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story deleted successfully!"})
}
