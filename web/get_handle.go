package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
)

func GetAllStories(context *gin.Context) {
	var stories []model.Story
	data.DB.Find(&stories)

	var dtos []model.StoryDto
	for _, item := range stories {
		dtos = append(dtos, model.StoryDto{ID: item.ID, Title: item.Title, Text: item.Text})
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

func GetStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")
	data.DB.Find(&story, id)

	dto := model.StoryDto{ID: story.ID, Title: story.Title, Text: story.Text}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}
