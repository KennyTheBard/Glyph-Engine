package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
	util "../util"
)

// CreateItem creates an item
func CreateItem(context *gin.Context) {
	var item model.ItemModel
	if err := context.BindJSON(&item); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	item, err := data.SaveItem(item)
	if err != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Failed to create new item!")
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Item created successfully!", "resourceId": item.ID})
}

// GetAllItems retrieves all items
func GetAllItems(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": data.FindAllItems()})
}

// GetItem retrieves an item
func GetItem(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	item, err := data.FindItemById(uint(id))
	if err != nil {
		util.StatusResponse(context, http.StatusNotFound, "No item for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": item.ToDto()})
}

// UpdateItem updates an item
func UpdateItem(context *gin.Context) {
	var updateItem model.ItemModel
	if err := context.BindJSON(&updateItem); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}

	data.UpdateItemField(uint(id), map[string]interface{}{
		"name": updateItem.Name,
		"text": updateItem.Text,
	})

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "item updated successfully!"})
}

// DeleteItem removes an item
func DeleteItem(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "id parameter is not an unsigned integer!")
		return
	}
	err = data.DeleteItemById(uint(id))
	if err != nil {
		util.StatusResponse(context, http.StatusNotFound, "No item for the given ID!")
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item deleted successfully!"})
}
