package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var userService *Service

type UserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Endpoint(db *sql.DB, rg *gin.RouterGroup) {
	userService = NewService(db)

	rg.POST("/register", registerEndpoint)
	rg.POST("/login", loginEndpoint)
}

func registerEndpoint(ctx *gin.Context) {
	var dto UserDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(dto.Username) < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username must have at least 8 characters"})
		return
	}

	if len(dto.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password must have at least 8 characters"})
		return
	}

	userService.Register(dto.Username, dto.Password)

	ctx.JSON(http.StatusCreated, gin.H{"status": "Successfully registered"})
}

func loginEndpoint(ctx *gin.Context) {
	var dto UserDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService.Login(dto.Username, dto.Password)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully logged in"})
}
