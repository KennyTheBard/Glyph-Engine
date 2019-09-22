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

func UpdateItemField(id uint, fields map[string]interface{}) error {
	item, err := FindItemById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&item).Update(name, value)
	}

	return nil
}

func DeleteItemById(id uint) error {
	var item model.ItemModel
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&item, id)
	if item.ID != id {
		return errors.New("No item found with the given ID")
	}
	DB.Delete(&item)
	return nil
}
