package main

import (
	"github.com/gin-gonic/gin"

	data "./data"
	web "./web"
)

var router *gin.Engine

func main() {
	data.Init()

	router = gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", web.GetAllStories)
		v1.GET("/:id", web.GetStory)
	}

	router.Run()
}
