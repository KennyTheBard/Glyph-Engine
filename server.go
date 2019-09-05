package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.GET("/api", handleGet)
	router.Run()
}

func handleGet(context *gin.Context) {
	msg, _ := context.GetQuery("name")
	context.String(http.StatusOK, "Your name is "+msg)
}
