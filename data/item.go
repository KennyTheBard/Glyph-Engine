package data

import "errors"

// Item is the principal form of resource in the game
type ItemModel struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type`
}

func (item ItemModel) ToDto() (ret struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type`
}) {
	ret.ID = item.ID
	ret.Name = item.Name
	ret.Text = item.Text
	ret.Type = item.Type

	return
}

// CRUD methods

func (item *ItemModel) Save() error {
	DB.Save(item)
	return nil
}

func (item *ItemModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(item, id)
	if item.ID != id {
		return errors.New("No item found with the given ID")
	}

	return nil
}

func (item *ItemModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(item).Update(fieldName, fieldValue)

	return nil
}

func (item *ItemModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(item).Update(name, value)
	}

	return nil
}

func (item *ItemModel) Delete() error {
	DB.Delete(item)
	return nil
}
