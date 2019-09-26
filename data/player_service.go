package data

import (
	"errors"

	model "../model"
)

func SavePlayer(player model.PlayerModel) (model.PlayerModel, error) {
	DB.Save(&player)
	return player, nil
}

func FindPlayerById(id uint) (model.PlayerModel, error) {
	var player model.PlayerModel
	if id == 0 {
		return player, errors.New("ID's must be positive numbers")
	}

	DB.First(&player, id)
	if player.ID != id {
		return player, errors.New("No player found with the given ID")
	}
	return player, nil
}

func FindPlayerByUsername(username string) (model.PlayerModel, error) {
	var player model.PlayerModel

	DB.Where("username = ?", username).First(&player)
	if player.Username != username {
		return player, errors.New("No such user found")
	}
	return player, nil
}

func FindAllPlayers() []model.PlayerModel {
	var player []model.PlayerModel
	DB.Find(&player)
	return player
}

func UpdatePlayerField(player model.PlayerModel, fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(&player).Update(name, value)
	}

	return nil
}

func DeletePlayer(player model.PlayerModel) error {
	DB.Delete(&player)
	return nil
}
