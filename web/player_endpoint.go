package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../data"
	util "../util"
)

func SignIn(context *gin.Context) {
	var player data.PlayerModel
	if err := context.BindJSON(&player); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	var aux data.PlayerModel
	if aux.FindByUsername(player.Username) == nil {
		util.StatusResponse(context, http.StatusConflict, "Username already in use!")
		return
	}

	// TODO: add account validator through wrapper function
	var err error
	player.Password, err = util.HashPassword(player.Password)

	if err != nil {
		util.StatusResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	player.Save()
	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Account created successfully!"})
}

func LogIn(context *gin.Context) {
	var logInData data.PlayerModel
	if err := context.BindJSON(&logInData); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	var player data.PlayerModel
	if player.FindByUsername(logInData.Username) != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Wrong username or password!")
		return
	}

	if util.CheckPasswordHash(logInData.Password, player.Password) {
		// TODO: return a session token
		context.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "Logged into your account successfully!"})
	} else {
		context.JSON(http.StatusCreated, gin.H{"status": http.StatusBadRequest, "message": "Wrong username or password!"})
	}
}
