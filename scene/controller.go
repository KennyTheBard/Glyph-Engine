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
	Title string `json:"title"`
	Text  string `form:"text"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	sceneService = NewService(db)

	rg.POST("/:storyId", createScene)
	rg.GET("/", getAllScenes)
	rg.GET("/:id", getScene)
	rg.PUT("/:id", updateScene)
	rg.DELETE("/:id", deleteScene)
}

func createScene(ctx *gin.Context) {
	storyId, err := strconv.Atoi(ctx.Param("storyId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dto SceneDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sceneService.Create(dto.Title, dto.Text, storyId)

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllScenes(ctx *gin.Context) {
	all, err := sceneService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func getScene(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := sceneService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func updateScene(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dto SceneDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = sceneService.Update(id, dto.Title, dto.Text); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteScene(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = sceneService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
