package data

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// AttributeStack is a pair of an stack and a number of said items
type AttributeStack struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Number      uint   `json:"number"`
	StackType   string `json:"stackType" 		gorm:"column:stack_type`
	AttributeID uint   `json:"attributeID" 		gorm:"column:attribute_id"`
	OwnerID     uint   `json:"ownerID" 		gorm:"column:owner_id"`
	OwnerType   string `json:"ownerType" 		gorm:"column:owner_type"`
}

func (stack AttributeStack) ToDto() gin.H {
	ret := make(gin.H)
	ret["id"] = stack.ID
	ret["number"] = stack.Number
	ret["stackType"] = stack.StackType
	ret["attributeID"] = stack.AttributeID
	ret["ownerID"] = stack.OwnerID
	ret["ownerType"] = stack.OwnerType

	return ret
}

func (stack *AttributeStack) GetAttribute() (AttributeModel, error) {
	var attribute AttributeModel
	err := attribute.FindById(stack.AttributeID)
	return attribute, err
}

// CRUD methods

func (stack *AttributeStack) Save() error {
	DB.Save(stack)
	return nil
}

func (stack *AttributeStack) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(stack, id)
	if stack.ID != id {
		return errors.New("No stack found with the given ID")
	}

	return nil
}

func (stack *AttributeStack) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(stack).Update(fieldName, fieldValue)

	return nil
}

func (stack *AttributeStack) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(stack).Update(name, value)
	}

	return nil
}

func (stack *AttributeStack) Delete() error {
	DB.Delete(stack)
	return nil
}
