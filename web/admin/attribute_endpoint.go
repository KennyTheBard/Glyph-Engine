package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../../data"
	util "../../util"
	validator "../../validator"
)

// CreateAttribute creates an attribute
func CreateAttribute(context *gin.Context) {
	var attribute data.AttributeModel
	if err := context.BindJSON(&attribute); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	if err := validator.Validate(attribute); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if attribute.Save() != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new attribute!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Attribute created successfully!", "resourceId": attribute.ID})
}

// GetAttribute retrieves an attribute
func GetAttribute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var attribute data.AttributeModel
	if attribute.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No attribute for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": attribute.ToDto()})
}

// UpdateAttribute updates an attribute
func UpdateAttribute(context *gin.Context) {
	var updateAttribute data.AttributeModel
	if err := context.BindJSON(&updateAttribute); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	if err := validator.Validate(updateAttribute); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var attribute data.AttributeModel
	if attribute.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No attribute for the given ID!")
		return
	}

	if err := attribute.UpdateFields(map[string]interface{}{
		"name": updateAttribute.Name,
		"text": updateAttribute.Text,
	}); err != nil {
		util.StatusResponse(context, http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute updated successfully!"})
}

// DeleteAttribute removes an attribute
func DeleteAttribute(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var attribute data.AttributeModel
	if attribute.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No attribute for the given ID!")
		return
	}

	if err := attribute.Delete(); err != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute deleted successfully!"})
}
