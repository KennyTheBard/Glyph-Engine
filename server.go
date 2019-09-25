package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	data "./data"
	web "./web"
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
			storyEndpoint.GET("/:id/choice", web.GetStoryChoices)
			storyEndpoint.PUT("/:id", web.UpdateStory)
			storyEndpoint.PUT("/:id/choice/:choiceid", web.AddChoiceToStory)
			storyEndpoint.DELETE("/:id", web.DeleteStory)
		}
		choiceEndpoint := api.Group("/choice")
		{
			choiceEndpoint.POST("/", web.CreateChoice)
			choiceEndpoint.GET("/", web.GetAllChoices)
			choiceEndpoint.GET("/:id", web.GetChoice)
			// choiceEndpoint.GET("/:id/cost", web.GetChoiceCosts)
			// choiceEndpoint.GET("/:id/reward", web.GetChoiceRewards)
			// choiceEndpoint.GET("/:id/requirement", web.GetChoiceRewards)
			choiceEndpoint.PUT("/:id", web.UpdateChoice)
			// choiceEndpoint.PUT("/:id/cost", web.UpdateChoiceCosts)
			// choiceEndpoint.PUT("/:id/reward", web.UpdateChoiceRewards)
			// choiceEndpoint.PUT("/:id/requirement", web.UpdateChoiceRequirements)
			choiceEndpoint.DELETE("/:id", web.DeleteChoice)
		}
		itemEndpoint := api.Group("/item")
		{
			itemEndpoint.POST("/", web.CreateItem)
			itemEndpoint.GET("/", web.GetAllItems)
			itemEndpoint.GET("/:id", web.GetItem)
			itemEndpoint.PUT("/:id", web.UpdateItem)
			itemEndpoint.DELETE("/:id", web.DeleteItem)
		}
	}

	if err := router.Run(); err != nil {
		panic(err)
	}

	// stories := data.FindAllStories()
	// intStories := make([]interface{}, len(stories))
	// for i, story := range stories {
	// 	intStories[i] = story
	// }

	// port.Export("test_export.txt", intStories)
	// port.Import("test_export.txt", 100, func(bs []byte) {
	// 	fmt.Println(string(bs))
	// })
}
