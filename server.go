package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	data "./data"
)

var router *gin.Engine

func main() {
	cleanStart := flag.Bool("cleanStart", false, "If the connection to the database should make sure the database is empty")
	flag.Parse()

	data.Init(*cleanStart)
	defer data.Close()

	router = gin.Default()

	api := router.Group("/api")
	{
		storyEndpoint := api.Group("/story")
		{
			storyEndpoint.POST("/", web.CreateStory)
			storyEndpoint.GET("/", web.GetAllStories)
			storyEndpoint.GET("/:id", web.GetStory)
			// storyEndpoint.GET("/:id/choice", web.GetStoryChoices)
			storyEndpoint.PUT("/:id", web.UpdateStory)
			// storyEndpoint.PUT("/:id/choice", web.UpdateStoryChoices)
			storyEndpoint.DELETE("/:id", web.DeleteStory)
			// storyEndpoint.DELETE("/:id/choice", web.DeleteStoryChoice)
		}
		choiceEndpoint := api.Group("/choice")
		{
			choiceEndpoint.POST("/", web.CreateChoice)
			choiceEndpoint.GET("/", web.GetAllChoices)
			choiceEndpoint.GET("/:id", web.GetChoice)
			choiceEndpoint.PUT("/:id", web.UpdateChoice)
			choiceEndpoint.DELETE("/:id", web.DeleteChoice)
		}
	}

	if err := router.Run(); err != nil {
		panic(err)
	}
}
