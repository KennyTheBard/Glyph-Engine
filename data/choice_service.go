package data

import (
	"errors"

	model "../model"
)

func SaveChoice(choice model.Choice) (model.Choice, error) {
	DB.Save(&choice)
	return choice, nil
}

func FindChoiceById(id uint) (model.Choice, error) {
	var choice model.Choice
	if id == 0 {
		return choice, errors.New("ID's must be positive numbers")
	}

	DB.First(&choice, id)
	if choice.ID != id {
		return choice, errors.New("No choice found with the given ID")
	}
	return choice, nil
}

func FindAllChoices() []model.Choice {
	var choices []model.Choice
	DB.Find(&choices)
	return choices
}

func UpdateChoiceField(id uint, fields map[string]interface{}) error {
	choice, err := FindChoiceById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&choice).Update(name, value)
	}

	return nil
}

func DeleteChoiceById(id uint) error {
	var choice model.Choice
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&choice, id)
	if choice.ID != id {
		return errors.New("No choice found with the given ID")
	}
	DB.Delete(&choice)
	return nil
}
