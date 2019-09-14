package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
)

func HandleGet(context *gin.Context) {
	var stories []model.Story
	var dtos []model.StoryDto

	data.DB.Find(&stories)

	if len(stories) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No story found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range stories {
		dtos = append(dtos, model.StoryDto{ID: item.ID, Title: item.Title, Text: item.Text})
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dtos})
}
