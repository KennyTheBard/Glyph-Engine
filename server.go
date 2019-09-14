package main

import (
	"github.com/gin-gonic/gin"

	data "./data"
	web "./web"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", web.HandleGet)
	}

	data.Init()

	router.Run()
}
