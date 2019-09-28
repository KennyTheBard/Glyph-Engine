package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	data "./data"
	// timeline "./timeline"
	web "./web"
)

var router *gin.Engine

func main() {
	cleanStart := flag.Bool("cleanStart", false, "clean init the database")
	flag.Parse()

	data.Init(*cleanStart)
	defer data.Close()

	// tm := timeline.NewTimeMachine()
	// tm.Start()

	// duration, _ := time.ParseDuration("10s")
	// tm.AddTimePoint(timeline.TimePoint{
	// 	Point:    time.Now().Add(duration),
	// 	WaitTime: duration,
	// 	IsRepetable: func() bool {
	// 		return true
	// 	},
	// 	Action: func() {
	// 		fmt.Println("Hello there!")
	// 	},
	// })

	router = gin.Default()

	api := router.Group("/api")
	{
		storyEndpoint := api.Group("/story")
		{
			storyEndpoint.POST("/", web.CreateStory)
			storyEndpoint.GET("/:id", web.GetStory)
			storyEndpoint.GET("/:id/choice", web.GetStoryChoices)
			storyEndpoint.PUT("/:id", web.UpdateStory)
			storyEndpoint.PUT("/:id/choice/:choiceid", web.AddChoiceToStory)
			storyEndpoint.DELETE("/:id", web.DeleteStory)
		}
		choiceEndpoint := api.Group("/choice")
		{
			choiceEndpoint.POST("/", web.CreateChoice)
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
			itemEndpoint.GET("/:id", web.GetItem)
			itemEndpoint.PUT("/:id", web.UpdateItem)
			itemEndpoint.DELETE("/:id", web.DeleteItem)
		}
		accountEndpoint := api.Group("/account")
		{
			accountEndpoint.POST("/sign", web.SignIn)
			accountEndpoint.POST("/log", web.LogIn)
			// accountEndpoint.PUT("/:id", web.UpdateAccount)
			// accountEndpoint.DELETE("/:id", web.DeactivateAccount)
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

	// tm.Stop()

}
