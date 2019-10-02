package data

import "errors"

const (
	OWNER_CHOICE = 1
	OWNER_USER   = 2
)

// ItemStack is a pair of an stack and a number of said items
type ItemStack struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Number    uint   `json:"number"`
	StackType string `json:"stackType" 		gorm:"column:stack_type`
	ItemID    uint   `json:"itemID" 		gorm:"column:item_id"`
	OwnerID   uint   `json:"ownerID" 		gorm:"column:owner_id"`
	OwnerType string `json:"ownerType" 		gorm:"column:owner_type"`
}

func (stack ItemStack) ToDto() (ret struct {
	ID        uint   `json:"id"`
	Number    uint   `json:"number"`
	StackType string `json:"stackType"`
	ItemID    uint   `json:"itemID"`
	OwnerID   uint   `json:"ownerID"`
	OwnerType string `json:"ownerType`
}) {
	ret.ID = stack.ID
	ret.Number = stack.Number
	ret.StackType = stack.StackType
	ret.ItemID = stack.ItemID
	ret.OwnerID = stack.OwnerID
	ret.OwnerType = stack.OwnerType

	return
}

func (stack *ItemStack) GetItem() (ItemModel, error) {
	var item ItemModel
	err := item.FindById(stack.ItemID)
	return item, err
}

// CRUD methods

func (stack *ItemStack) Save() error {
	DB.Save(stack)
	return nil
}

func (stack *ItemStack) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(stack, id)
	if stack.ID != id {
		return errors.New("No stack found with the given ID")
	}

	return nil
}

func (stack *ItemStack) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(stack).Update(fieldName, fieldValue)

	return nil
}

func (stack *ItemStack) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(stack).Update(name, value)
	}

	return nil
}

func (stack *ItemStack) Delete() error {
	DB.Delete(stack)
	return nil
}
