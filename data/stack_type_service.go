package data

import (
	"errors"

	model "../model"
)

func SaveStackType(stackType model.StackType) (model.StackType, error) {
	DB.Save(&stackType)
	return stackType, nil
}

func FindStackTypeById(id uint) (model.StackType, error) {
	var stackType model.StackType
	if id == 0 {
		return stackType, errors.New("ID's must be positive numbers")
	}

	DB.First(&stackType, id)
	if stackType.ID != id {
		return stackType, errors.New("No stack type found with the given ID")
	}
	return stackType, nil
}

func FindAllStackTypes() []model.StackType {
	var stackTypes []model.StackType
	DB.Find(&stackTypes)
	return stackTypes
}

func UpdateStackTypeField(id uint, fields map[string]interface{}) error {
	stackType, err := FindStackTypeById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&stackType).Update(name, value)
	}

	return nil
}

func DeleteStackTypeById(id uint) error {
	var stackType model.StackType
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&stackType, id)
	if stackType.ID != id {
		return errors.New("No stackType found with the given ID")
	}
	DB.Delete(&stackType)
	return nil
}

func DeleteStackType(stackType model.StackType) error {
	DB.Delete(&stackType)
	return nil
}
