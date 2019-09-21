package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	data "./data"
	model "./model"
	// web "./web"
)

var router *gin.Engine

func main() {
	data.Init()
	defer data.Close()

	// router = gin.Default()

	// api := router.Group("/api")
	// {
	// 	storyEndpoint := api.Group("/story")
	// 	{
	// 		storyEndpoint.POST("/", web.CreateStory)
	// 		storyEndpoint.GET("/", web.GetAllStories)
	// 		storyEndpoint.GET("/:id", web.GetStory)
	// 		storyEndpoint.PUT("/:id", web.UpdateStory)
	// 		storyEndpoint.DELETE("/:id", web.DeleteStory)
	// 	}
	// 	choiceEndpoint := api.Group("/choice")
	// 	{
	// 		choiceEndpoint.POST("/", web.CreateChoice)
	// 		choiceEndpoint.GET("/", web.GetAllChoices)
	// 		choiceEndpoint.GET("/:id", web.GetChoice)
	// 		choiceEndpoint.PUT("/:id", web.UpdateChoice)
	// 		choiceEndpoint.DELETE("/:id", web.DeleteChoice)
	// 	}
	// }

	// if err := router.Run(); err != nil {
	// 	panic(err)
	// }

	data.SaveStory(model.Story{Title: "GOOD-STORY", Text: "etc"})
	fmt.Println(data.FindStoryById(1))
	story, _ := data.FindStoryById(1)
	fmt.Println(story.ToDto())
}
