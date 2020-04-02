package story

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var choiceService *Service

type ChoiceDTO struct {
	Title     string `json:"name"`
	Text      string `form:"text"`
	SceneId   int    `json:"scene_id" binding:"required"`
	NextScene int    `json:"next_scene" binding:"required"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	choiceService = NewService(db)

	rg.POST("/", createChoice)
	rg.GET("/", getAllChoices)
	rg.GET("/:id", getChoice)
	rg.PUT("/:id", updateChoice)
	rg.DELETE("/:id", deleteChoice)
}

func createChoice(ctx *gin.Context) {
	var dto ChoiceDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	choiceService.Create(dto.Title, dto.Text, dto.SceneId, dto.NextScene)

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllChoices(ctx *gin.Context) {
	all, err := choiceService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func getChoice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := choiceService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func updateChoice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dto ChoiceDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = choiceService.Update(id, dto.Title, dto.Text, dto.SceneId, dto.NextScene); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteChoice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = choiceService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
