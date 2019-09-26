package data

import (
	"errors"

	model "../model"
)

func SaveItem(item model.ItemModel) (model.ItemModel, error) {
	DB.Save(&item)
	return item, nil
}

func FindItemById(id uint) (model.ItemModel, error) {
	var item model.ItemModel
	if id == 0 {
		return item, errors.New("ID's must be positive numbers")
	}

	DB.First(&item, id)
	if item.ID != id {
		return item, errors.New("No item found with the given ID")
	}
	return item, nil
}

func FindAllItems() []model.ItemModel {
	var items []model.ItemModel
	DB.Find(&items)
	return items
}

func UpdateItemField(item model.ItemModel, fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(&item).Update(name, value)
	}

	return nil
}

func DeleteItem(item model.ItemModel) error {
	DB.Delete(&item)
	return nil
}
