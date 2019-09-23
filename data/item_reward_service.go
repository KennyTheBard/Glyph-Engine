package data

import (
	"errors"

	model "../model"
)

func SaveItemReward(reward model.ItemReward) (model.ItemReward, error) {
	DB.Save(&reward)
	return reward, nil
}

func FindItemRewardById(id uint) (model.ItemReward, error) {
	var reward model.ItemReward
	if id == 0 {
		return reward, errors.New("ID's must be positive numbers")
	}

	DB.First(&reward, id)
	if reward.ID != id {
		return reward, errors.New("No reward found with the given ID")
	}
	return reward, nil
}

func FindAllItemRewards() []model.ItemReward {
	var rewards []model.ItemReward
	DB.Find(&rewards)
	return rewards
}

func UpdateItemRewardField(id uint, fields map[string]interface{}) error {
	reward, err := FindItemRewardById(id)
	if err != nil {
		return err
	}

	for name, value := range fields {
		DB.Model(&reward).Update(name, value)
	}

	return nil
}

func DeleteItemRewardById(id uint) error {
	var reward model.ItemReward
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(&reward, id)
	if item.ID != id {
		return errors.New("No reward found with the given ID")
	}
	DB.Delete(&reward)
	return nil
}
