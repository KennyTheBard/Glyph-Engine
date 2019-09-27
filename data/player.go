package data

import "errors"

// PlayerModel is the main element of a page
type PlayerModel struct {
	ID       uint   `json:"id" 		gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (player PlayerModel) ToDto() (ret struct {
	ID       uint   `json:"id" 		gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}) {
	ret.ID = player.ID
	ret.Username = player.Username
	ret.Password = player.Password

	return
}

func (player *PlayerModel) Save() error {
	DB.Save(player)
	return nil
}

func (player *PlayerModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(player, id)
	if player.ID != id {
		return errors.New("No player found with the given ID")
	}

	return nil
}

func (player *PlayerModel) FindByUsername(username string) error {
	DB.Where("username = ?", username).First(player)
	if player.Username != username {
		return errors.New("No player found with the given ID")
	}

	return nil
}

func (player *PlayerModel) GetInventory() []ItemStack {
	var stacks []ItemStack
	DB.Where("owner_id = ? and type = ?", player.ID, OWNER_PLAYER).Find(&stacks)
	return stacks
}

func (player *PlayerModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(player).Update(fieldName, fieldValue)

	return nil
}

func (player *PlayerModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(player).Update(name, value)
	}

	return nil
}

func (player *PlayerModel) Delete() error {
	DB.Delete(player)
	return nil
}
