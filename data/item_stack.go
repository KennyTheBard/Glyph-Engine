package data

import "errors"

const (
	OWNER_CHOICE = 1
	OWNER_PLAYER = 2
)

// ItemStack is a pair of an stack and a number of said items
type ItemStack struct {
	ID        uint   `json:"id" 			gorm:"primary_key"`
	Number    uint   `json:"number"`
	Type      string `json:"Type" 			gorm:"column:type`
	ItemID    uint   `json:"itemID" 		gorm:"column:item_id"`
	OwnerID   uint   `json:"ownerID" 		gorm:"column:owner_id"`
	OwnerType uint   `json:"ownerType" 		gorm:"column:owner_type"`
}

func (stack ItemStack) ToDto() (ret struct {
	ID        uint   `json:"id"`
	Number    uint   `json:"number"`
	Type      string `json:"type"`
	ItemID    uint   `json:"itemID"`
	OwnerID   uint   `json:"ownerID"`
	OwnerType uint   `json:"ownerType`
}) {
	ret.ID = stack.ID
	ret.Number = stack.Number
	ret.Type = stack.Type
	ret.ItemID = stack.ItemID
	ret.OwnerID = stack.OwnerID
	ret.OwnerType = stack.OwnerType

	return
}

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
		return errors.New("No item stack found with the given ID")
	}

	return nil
}

func (stack *ItemStack) GetItem() ItemModel {
	var item ItemModel
	DB.First(&item, stack.ItemID)
	return item
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
