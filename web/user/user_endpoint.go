package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../../data"
	security "../../security"
	util "../../util"
)

func SignIn(context *gin.Context) {
	var user data.UserModel
	if err := context.BindJSON(&user); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	var aux data.UserModel
	if aux.FindByUsername(user.Username) == nil {
		util.StatusResponse(context, http.StatusConflict, "Username already in use!")
		return
	}

	// TODO: add account validator through wrapper function
	var err error
	user.Password, err = security.HashPassword(user.Password)

	if err != nil {
		util.StatusResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	// replace thid hardcoded value with configurable one
	user.CurrStoryID = 1
	user.UserType = "admin"

	if err = user.Save(); err != nil {
		util.StatusResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Account created successfully!"})
}

func LogIn(context *gin.Context) {
	var logInData data.UserModel
	if err := context.BindJSON(&logInData); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	var user data.UserModel
	if user.FindByUsername(logInData.Username) != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Wrong username or password!")
		return
	}

	if security.CheckPasswordHash(logInData.Password, user.Password) {
		context.JSON(http.StatusCreated, gin.H{
			"status":      http.StatusOK,
			"message":     "Logged into your account successfully!",
			"accessToken": security.Authorizate(user.Username),
		})
	} else {
		util.StatusResponse(context, http.StatusCreated, "Wrong username or password!")
	}
}
