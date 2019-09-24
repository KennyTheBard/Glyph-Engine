package data

import (
	"errors"

	model "../model"
)

func SaveItemStack(stack model.ItemStack) (model.ItemStack, error) {
	DB.Save(&stack)
	return stack, nil
}

func FindItemStackById(id uint) (model.ItemStack, error) {
	var stack model.ItemStack
	if id == 0 {
		return stack, errors.New("ID's must be positive numbers")
	}

	DB.First(&stack, id)
	if stack.ID != id {
		return stack, errors.New("No item stack found with the given ID")
	}
	return stack, nil
}

func FindAllItemStacks() []model.ItemStack {
	var stacks []model.ItemStack
	DB.Find(&stacks)
	return stacks
}

func UpdateItemStackField(id uint, fields map[string]interface{}) error {
	stack, err := FindItemStackById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&stack).Update(name, value)
	}

	return nil
}

func DeleteItemStackById(id uint) error {
	var stack model.ItemStack
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&stack, id)
	if stack.ID != id {
		return errors.New("No item stack found with the given ID")
	}
	DB.Delete(&stack)
	return nil
}

func DeleteItemStack(stack model.ItemStack) error {
	DB.Delete(&stack)
	return nil
}
