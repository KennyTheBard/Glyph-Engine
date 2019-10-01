package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../../data"
	security "../../security"
	util "../../util"
)

func GetCurrentStory(context *gin.Context) {
	usernames := context.Request.Header["Username"]
	tokens := context.Request.Header["Access-Token"]

	if len(usernames) != 1 || len(tokens) != 1 {
		util.StatusResponse(context, http.StatusBadRequest, "No credentials provided")
		return
	}

	username := usernames[0]
	token := tokens[0]

	if !security.VerifyToken(username, token) {
		util.StatusResponse(context, http.StatusForbidden, "Session expired!")
		return
	}

	var player data.PlayerModel
	if player.FindByUsername(username) != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Unknown username!")
		return
	}

	story := player.GetCurrentStory()

	// TODO: return story dtos
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"story":   story.ToDto(),
		"choices": story.GetChoices(),
	})
}

func MakeChoice(context *gin.Context) {
	usernames := context.Request.Header["Username"]
	tokens := context.Request.Header["Access-Token"]

	if len(usernames) != 1 || len(tokens) != 1 {
		util.StatusResponse(context, http.StatusBadRequest, "No credentials provided")
		return
	}

	username := usernames[0]
	token := tokens[0]

	if !security.VerifyToken(username, token) {
		util.StatusResponse(context, http.StatusForbidden, "Session expired!")
		return
	}

	var player data.PlayerModel
	if player.FindByUsername(username) != nil {
		util.StatusResponse(context, http.StatusInternalServerError, "Unknown username!")
		return
	}

	var recievedChoice data.ChoiceModel
	if err := context.BindJSON(&recievedChoice); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, "Missing or incorrect object sent!")
		return
	}

	var choice data.ChoiceModel
	if err := choice.FindById(recievedChoice.ID); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	story := player.GetCurrentStory()
	if choice.ParentStoryID != story.ID {
		util.StatusResponse(context, http.StatusForbidden, "You cannot choose this story")
		return
	}

	if err := player.UpdateField("curr_story_id", choice.NextStoryID); err != nil {
		util.StatusResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
