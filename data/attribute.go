package data

import "errors"

// attribute is the principal form of resource in the game
type AttributeModel struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type`
}

func (attribute AttributeModel) ToDto() (ret struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type`
}) {
	ret.ID = attribute.ID
	ret.Name = attribute.Name
	ret.Text = attribute.Text
	ret.Type = attribute.Type

	return
}

// CRUD methods

func (attribute *AttributeModel) Save() error {
	DB.Save(attribute)
	return nil
}

func (attribute *AttributeModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(attribute, id)
	if attribute.ID != id {
		return errors.New("No attribute found with the given ID")
	}

	return nil
}

func (attribute *AttributeModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(attribute).Update(fieldName, fieldValue)

	return nil
}

func (attribute *AttributeModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(attribute).Update(name, value)
	}

	return nil
}

func (attribute *AttributeModel) Delete() error {
	DB.Delete(attribute)
	return nil
}
