package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
)

// CreateStory creates a story
func CreateStory(context *gin.Context) {
	var story model.StoryModel
	if err := context.BindJSON(&story); err != nil {
		StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	story, err := data.SaveStory(story)
	if err != nil {
		StatusResponse(context, http.StatusInternalServerError, "Failed to create new story!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Story created successfully!", "resourceId": story.ID})
}

// GetAllStories retrieves all stories
func GetAllStories(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": data.FindAllStories()})
}

// GetStory retrieves a story
func GetStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	story, err := data.FindStoryById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": story.ToDto()})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var updatedStory model.StoryModel
	if err := context.BindJSON(&updatedStory); err != nil {
		StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	story, err := data.FindStoryById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	data.DB.Model(&story).Update("title", updatedStory.Title)
	data.DB.Model(&story).Update("text", updatedStory.Text)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story updated successfully!"})
}

// DeleteStory removes a story
func DeleteStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	err = data.DeleteStoryById(uint(id))
	if err != nil {
		StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	StatusResponse(context, http.StatusOK, "Story deleted successfully!")
}
