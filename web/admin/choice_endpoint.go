package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../../data"
	util "../../util"
)

// CreateChoice creates a choice
func CreateChoice(context *gin.Context) {
	var choice data.ChoiceModel
	if err := context.BindJSON(&choice); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	if choice.Save() != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new choice!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Choice created successfully!", "resourceId": choice.ID})
}

// GetChoice retrieves a choice
func GetChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var choice data.ChoiceModel
	if choice.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": choice.ToDto()})
}

// UpdateChoice updates a choice
func UpdateChoice(context *gin.Context) {
	var updateChoice data.ChoiceModel
	if err := context.BindJSON(&updateChoice); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var choice data.ChoiceModel
	if choice.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	if err := choice.UpdateFields(map[string]interface{}{
		"name":            updateChoice.Name,
		"text":            updateChoice.Text,
		"parent_story_id": updateChoice.ParentStoryID,
		"next_story_id":   updateChoice.NextStoryID,
	}); err != nil {
		util.StatusResponse(context, http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice updated successfully!"})
}

// DeleteChoice removes a choice
func DeleteChoice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var choice data.ChoiceModel
	if choice.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No choice for the given ID!")
		return
	}

	if err := choice.Delete(); err != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Choice deleted successfully!"})
}
