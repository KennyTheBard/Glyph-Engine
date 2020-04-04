package story

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var progressService *Service

type ProgressDTO struct {
	UserId  int `json:"user_id" binding:"required"`
	StoryId int `form:"story_id" binding:"required"`
	SceneId int `json:"scene_id"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	progressService = NewService(db)

	rg.POST("/", createProgress)
	rg.GET("/", getAllProgress)
	rg.GET("/user/:id", getProgressByUserId)
	rg.GET("/story/:id", getProgressByStoryId)
	rg.PUT("/", updateProgress)
	rg.DELETE("/", deleteProgress)
}

func createProgress(ctx *gin.Context) {
	var dto ProgressDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := progressService.Create(dto.UserId, dto.StoryId, dto.SceneId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllProgress(ctx *gin.Context) {
	all, err := progressService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func getProgressByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	all, err := progressService.GetAllByUserId(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func getProgressByStoryId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	all, err := progressService.GetAllByStoryId(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func updateProgress(ctx *gin.Context) {
	var dto ProgressDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := progressService.Update(dto.UserId, dto.StoryId, dto.SceneId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteProgress(ctx *gin.Context) {
	var dto ProgressDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := progressService.Delete(dto.UserId, dto.StoryId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
