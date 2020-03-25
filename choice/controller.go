package story

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var sceneService *Service

type ChoiceDTO struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Text      string `form:"text"`
	SceneId   int    `json:"scene_id" binding:"required"`
	NextScene int    `json:"next_scene" binding:"required"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	sceneService = NewService(db)

	rg.POST("/", createChoice)
	rg.GET("/", getAllChoices)
	rg.GET("/:id", getChoice)
	rg.PUT("/", updateChoice)
	rg.DELETE("/", deleteChoice)
}

func createChoice(ctx *gin.Context) {
	var dto SceneDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sceneService.Create(dto.Title, dto.Text, dto.StoryId)

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllChoices(ctx *gin.Context) {
	all := sceneService.GetAll()

	ctx.JSON(http.StatusOK, all)
}

func getChoice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := sceneService.GetById(id)

	ctx.JSON(http.StatusOK, dto)
}

func updateChoice(ctx *gin.Context) {
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

func deleteChoice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sceneService.Delete(id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
