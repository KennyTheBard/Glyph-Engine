package util

import "github.com/gin-gonic/gin"

func StatusResponse(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{"status": code, "message": message})
}
