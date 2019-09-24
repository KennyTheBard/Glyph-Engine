package data

import (
	"errors"

	model "../model"
)

func SaveItemCost(cost model.ItemCost) (model.ItemCost, error) {
	DB.Save(&cost)
	return cost, nil
}

func FindItemCostById(id uint) (model.ItemCost, error) {
	var cost model.ItemCost
	if id == 0 {
		return cost, errors.New("ID's must be positive numbers")
	}

	DB.First(&cost, id)
	if cost.ID != id {
		return cost, errors.New("No cost found with the given ID")
	}
	return cost, nil
}

func FindAllItemCosts() []model.ItemCost {
	var costs []model.ItemCost
	DB.Find(&costs)
	return costs
}

func UpdateItemCostField(id uint, fields map[string]interface{}) error {
	cost, err := FindItemCostById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&cost).Update(name, value)
	}

	return nil
}

func DeleteItemCostById(id uint) error {
	var cost model.ItemCost
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&cost, id)
	if item.ID != id {
		return errors.New("No cost found with the given ID")
	}
	DB.Delete(&cost)
	return nil
}

func DeleteItemCost(cost model.ItemCost) error {
	DB.Delete(&cost)
	return nil
}
