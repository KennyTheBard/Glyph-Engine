package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	config "./config"
	data "./data"
	security "./security"
	tline "./timeline"
	admin "./web/admin"
	user "./web/user"
)

var router *gin.Engine

func main() {
	cleanStart := flag.Bool("cleanStart", false, "clean init the database")
	configFile := flag.String("config", "", "path to the config file")
	flag.Parse()

	data.Init(*cleanStart)
	defer data.Close()

	if config.LoadConfig(*configFile) != nil {
		if len(*configFile) == 0 {
			fmt.Println("No configuration file was given, starting on default configurations")
		} else {
			fmt.Println("Failed to read given configuration file, starting on default configurations")
		}
	}

	tline.Timeline.Init()
	tline.Timeline.Start()
	defer tline.Timeline.Stop()

	// duration, _ := time.ParseDuration("1s")
	// tline.Timeline.AddTimePoint(tline.TimePoint{
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
	router.Use(security.CORSMiddleware())

	api := router.Group("/api")
	{
		adminGroup := api.Group("/admin")
		{
			storyEndpoint := adminGroup.Group("/story")
			{
				storyEndpoint.POST("/", admin.CreateStory)
				storyEndpoint.GET("/:id", admin.GetStory)
				storyEndpoint.GET("/:id/choice", admin.GetStoryChoices)
				storyEndpoint.PUT("/:id", admin.UpdateStory)
				storyEndpoint.PUT("/:id/choice/:choiceid", admin.AddChoiceToStory)
				storyEndpoint.DELETE("/:id", admin.DeleteStory)
			}
			choiceEndpoint := adminGroup.Group("/choice")
			{
				choiceEndpoint.POST("/", admin.CreateChoice)
				choiceEndpoint.GET("/:id", admin.GetChoice)
				// choiceEndpoint.GET("/:id/cost", admin.GetChoiceCosts)
				// choiceEndpoint.GET("/:id/reward", admin.GetChoiceRewards)
				// choiceEndpoint.GET("/:id/requirement", admin.GetChoiceRewards)
				choiceEndpoint.PUT("/:id", admin.UpdateChoice)
				// choiceEndpoint.PUT("/:id/cost", admin.UpdateChoiceCosts)
				// choiceEndpoint.PUT("/:id/reward", admin.UpdateChoiceRewards)
				// choiceEndpoint.PUT("/:id/requirement", admin.UpdateChoiceRequirements)
				choiceEndpoint.DELETE("/:id", admin.DeleteChoice)
			}
			attributeEndpoint := adminGroup.Group("/attribute")
			{
				attributeEndpoint.POST("/", admin.CreateAttribute)
				attributeEndpoint.GET("/:id", admin.GetAttribute)
				attributeEndpoint.PUT("/:id", admin.UpdateAttribute)
				attributeEndpoint.DELETE("/:id", admin.DeleteAttribute)
			}
		}

		userGroup := api.Group("/user")
		{
			accountEndpoint := userGroup.Group("/account")
			{
				accountEndpoint.POST("/sign", user.SignIn)
				accountEndpoint.POST("/log", user.LogIn)
				// accountEndpoint.PUT("/:id", user.UpdateAccount)
				// accountEndpoint.DELETE("/:id", user.DeactivateAccount)
			}
			gameEndpoint := userGroup.Group("/game")
			{
				gameEndpoint.GET("/", user.GetCurrentStory)
				gameEndpoint.POST("/", user.MakeChoice)
			}
		}
	}

	if err := router.Run(":" + strconv.Itoa(int(config.GlobalConfiguration.Port))); err != nil {
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
