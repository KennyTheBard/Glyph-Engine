package data

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ChoiceModel is the main subelement of the page
type ChoiceModel struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	Text          string `json:"text"`
	ParentStoryID uint   `json:"parentStoryID" 	gorm:"column:parent_story_id"`
	ChoiceScript  string `json:"choiceScript" 		gorm:"column:choice_script"`
}

// DTO methods

func (choice ChoiceModel) ToDto() gin.H {
	ret := make(gin.H)
	ret["id"] = choice.ID
	ret["name"] = choice.Name
	ret["text"] = choice.Text
	ret["parentStoryID"] = choice.ParentStoryID
	ret["choiceScript"] = choice.ChoiceScript

	return ret
}

// Useful methods

func (choice *ChoiceModel) GetAttributeStacks() []AttributeStack {
	var stacks []AttributeStack
	DB.Where("owner_id = ? and type = ?", choice.ID, "choice").Find(&stacks)
	return stacks
}

func (choice *ChoiceModel) GetAttribute(attribute_name, stackType string) AttributeStack {
	var stack AttributeStack
	DB.Joins("JOIN attribute_model ON attribute_models.id = stacks.attribute_id AND attribute_models.name = ?", attribute_name).Where("owner_id = ? and owner_type = ? and stack_type = ?", choice.ID, "choice", stackType).First(&stack)
	return stack
}

// CRUD methods

func (choice *ChoiceModel) Save() error {
	DB.Save(choice)
	return nil
}

func (choice *ChoiceModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(choice, id)
	if choice.ID != id {
		return errors.New("No choice found with the given ID")
	}

	return nil
}

func (choice *ChoiceModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(choice).Update(fieldName, fieldValue)

	return nil
}

func (choice *ChoiceModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(choice).Update(name, value)
	}

	return nil
}

func (choice *ChoiceModel) Delete() error {
	DB.Delete(choice)
	return nil
}
