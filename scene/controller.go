package story

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var sceneService *Service

type SceneDTO struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Text    string `form:"text"`
	StoryId int    `json:"story_id" binding:"required"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	sceneService = NewService(db)

	rg.POST("/", createScene)
	rg.GET("/", getAllScenes)
	rg.GET("/:id", getScene)
	rg.PUT("/", updateScene)
	rg.DELETE("/", deleteScene)
}

func createScene(ctx *gin.Context) {
	var dto SceneDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sceneService.Create(dto.Title, dto.Text, dto.StoryId)

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllStories(ctx *gin.Context) {
	all := sceneService.GetAll()

	ctx.JSON(http.StatusOK, all)
}

func getStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := sceneService.GetById(id)

	ctx.JSON(http.StatusOK, dto)
}

func updateStory(ctx *gin.Context) {
	var dto StoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dto.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Provide id of the entity to be updated"})
		return
	}

	sceneService.Update(dto.Id, dto.Title, dto.Text, dto.StoryId)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sceneService.Delete(id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
