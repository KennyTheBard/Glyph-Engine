package data

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ResolutionModel is the main element of a page
type ResolutionModel struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// DTO methods

func (resolution ResolutionModel) ToDto() gin.H {
	ret := make(gin.H)
	ret["id"] = resolution.ID
	ret["name"] = resolution.Name
	ret["text"] = resolution.Text

	return ret
}

// Useful methods

func (resolution *AttributeStack) GetAttributeStacks() []AttributeStack {
	var stacks []AttributeStack
	DB.Where("owner_id = ? and type = ?", resolution.ID, "resolution").Find(&stacks)
	return stacks
}

// CRUD methods

func (resolution *ResolutionModel) Save() error {
	DB.Save(resolution)
	return nil
}

func (resolution *ResolutionModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(resolution, id)
	if resolution.ID != id {
		return errors.New("No resolution found with the given ID")
	}

	return nil
}

func (resolution *ResolutionModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(resolution).Update(fieldName, fieldValue)

	return nil
}

func (resolution *ResolutionModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(resolution).Update(name, value)
	}

	return nil
}

func (resolution *ResolutionModel) Delete() error {
	DB.Delete(resolution)
	return nil
}
