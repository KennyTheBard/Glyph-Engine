package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../../data"
	util "../../util"
	validator "../../validator"
)

// CreateItem creates an item
func CreateItem(context *gin.Context) {
	var item data.ItemModel
	if err := context.BindJSON(&item); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	if err := validator.Validate(item); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if item.Save() != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new item!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Item created successfully!", "resourceId": item.ID})
}

// GetItem retrieves an item
func GetItem(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var item data.ItemModel
	if item.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No item for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": item.ToDto()})
}

// UpdateItem updates an item
func UpdateItem(context *gin.Context) {
	var updateItem data.ItemModel
	if err := context.BindJSON(&updateItem); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	if err := validator.Validate(updateItem); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var item data.ItemModel
	if item.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No item for the given ID!")
		return
	}

	if err := item.UpdateFields(map[string]interface{}{
		"name": updateItem.Name,
		"text": updateItem.Text,
	}); err != nil {
		util.StatusResponse(context, http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "item updated successfully!"})
}

// DeleteItem removes an item
func DeleteItem(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	var item data.ItemModel
	if item.FindById(uint(id)) != nil {
		util.StatusResponse(context, http.StatusNotFound, "No item for the given ID!")
		return
	}

	if err := item.Delete(); err != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item deleted successfully!"})
}
