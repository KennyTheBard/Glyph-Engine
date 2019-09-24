package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
	util "../util"
)

// CreateStory creates a story
func CreateStory(context *gin.Context) {
	var story model.StoryModel
	if err := context.BindJSON(&story); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	story, err := data.SaveStory(story)
	if err != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new story!")
		return
	}

	for _, choice := range story.Choices {
		_, err := data.SaveChoice(choice)
		if err != nil {
			util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new choice!")
			return
		}
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
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	story, err := data.FindStoryById(uint(id))
	if err != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": story.ToDto()})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var updatedStory model.StoryModel
	if err := context.BindJSON(&updatedStory); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	if err := data.UpdateStoryField(uint(id), map[string]interface{}{
		"name": updatedStory.Name,
		"text": updatedStory.Text,
	}); err != nil {
		util.StatusResponse(context, http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story updated successfully!"})
}

// AddChoiceToStory add a choice to a story
func AddChoiceToStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	story, err := data.FindStoryById(uint(id))
	if err != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	choiceId, err := strconv.Atoi(context.Param("choiceid"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "choiceid parameter is not an unsigned integer!")
		return
	}

	choice, err := data.FindChoiceById(uint(choiceId))
	if err != nil {
		util.StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	choice.ParentStoryID = story.ID
	data.SaveChoice(choice)
	if err != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new choice!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Choice added to the story successfully!", "resourceId": choice.ID})
}

// DeleteStory removes a story
func DeleteStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	err = data.DeleteStoryById(uint(id))
	if err != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	util.StatusResponse(context, http.StatusOK, "Story deleted successfully!")
}
