package data

import (
	"errors"

	model "../model"
)

func SaveItemRequirement(requirement model.ItemRequirement) (model.ItemRequirement, error) {
	DB.Save(&requirement)
	return requirement, nil
}

func FindItemRequirementById(id uint) (model.ItemRequirement, error) {
	var requirement model.ItemRequirement
	if id == 0 {
		return requirement, errors.New("ID's must be positive numbers")
	}

	DB.First(&requirement, id)
	if requirement.ID != id {
		return requirement, errors.New("No requirement found with the given ID")
	}
	return requirement, nil
}

func FindAllItemRequirements() []model.ItemRequirement {
	var requirements []model.ItemRequirement
	DB.Find(&requirements)
	return requirements
}

func UpdateItemRequirementField(id uint, fields map[string]interface{}) error {
	requirement, err := FindItemRequirementById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&requirement).Update(name, value)
	}

	return nil
}

func DeleteItemRequirementById(id uint) error {
	var requirement model.ItemRequirement
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&requirement, id)
	if item.ID != id {
		return errors.New("No requirement found with the given ID")
	}
	DB.Delete(&requirement)
	return nil
}
