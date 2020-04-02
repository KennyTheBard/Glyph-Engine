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
	Title       string `json:"title" binding:"required"`
	Description string `form:"description"`
	AuthorId    int    `json:"author_id" binding:"required"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	storyService = NewService(db)

	rg.POST("/", createStory)
	rg.GET("/", getAllStories)
	rg.GET("/:id", getStory)
	rg.PUT("/:id", updateStory)
	rg.DELETE("/:id", deleteStory)
}

func createStory(ctx *gin.Context) {
	var dto StoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := storyService.Create(dto.Title, dto.Description, dto.AuthorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully created"})
}

func getAllStories(ctx *gin.Context) {
	all, err := storyService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func getStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := storyService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func updateStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Provide id of the entity to be updated"})
		return
	}

	var dto StoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = storyService.Update(id, dto.Title, dto.Description, dto.AuthorId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteStory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = storyService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
