package story

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var storyService *Service

type StoryDTO struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `form:"description"`
	AuthorId    int    `json:"author_id" binding:"required"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	storyService = NewService(db)

	rg.POST("/", createStory)
	rg.GET("/", getAllStories)
	rg.GET("/:id", getStory)
	rg.PUT("/", updateStory)
	rg.DELETE("/", deleteStory)
}

func createStory(ctx *gin.Context) {
	var dto StoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyService.Create(dto.Title, dto.Description, dto.AuthorId)

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllStories(ctx *gin.Context) {
	all := storyService.GetAll()

	ctx.JSON(http.StatusOK, all)
}

func getStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := storyService.GetById(id)

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

	storyService.Update(dto.Id, dto.Title, dto.Description, dto.AuthorId)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyService.Delete(id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
