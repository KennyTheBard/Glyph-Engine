package service

import (
	"errors"

	data "../data"
	model "../model"
)

func SaveChoice(choice model.Choice) (model.Choice, error) {
	data.DB.Save(&choice)
	return choice, nil
}

func FindChoiceById(id uint) (model.Choice, error) {
	var choice model.Choice
	data.DB.First(&choice, id)
	if choice.ID != id {
		return choice, errors.New("No choice found with the given ID")
	}
	return choice, nil
}

func FindAllChoices() []model.Choice {
	var choices []model.Choice
	data.DB.Find(&choices)
	return choices
}

func DeleteChoiceById(id uint) error {
	var choice model.Choice
	data.DB.First(&choice, id)
	if choice.ID != id {
		return errors.New("No choice found with the given ID")
	}
	data.DB.Delete(&choice)
	return nil
}
