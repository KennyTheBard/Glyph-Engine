package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../data"
	util "../util"
)

// CreateStory creates a story
func CreateStory(context *gin.Context) {
	var story data.StoryModel
	if err := context.BindJSON(&story); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	if story.Save() != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new story!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Story created successfully!", "resourceId": story.ID})
}

// GetStory retrieves a story
func GetStory(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var story data.StoryModel
	if story.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": story.ToDto()})
}

// GetStoryChoices retrives the choices of the story with the given id
func GetStoryChoices(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var story data.StoryModel
	if story.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	choices := story.GetChoices()
	choiceDtos := make([]interface{}, len(choices))
	for i, choice := range choices {
		choiceDtos[i] = choice.ToDto()
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": choiceDtos})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var updatedStory data.StoryModel
	if err := context.BindJSON(&updatedStory); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var story data.StoryModel
	if story.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	if err := story.UpdateFields(map[string]interface{}{
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

	var story data.StoryModel
	if story.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	choiceId, err := strconv.Atoi(context.Param("choiceid"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "choiceid parameter is not an unsigned integer!")
		return
	}

	var choice data.ChoiceModel
	if choice.FindById(uint(choiceId)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	choice.ParentStoryID = story.ID
	choice.Save()
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

	var story data.StoryModel
	if story.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No story for the given ID!")
		return
	}

	story.Delete()
	util.StatusResponse(context, http.StatusOK, "Story deleted successfully!")
}
