package model

// ItemStack is a pair of an item and a number of said items
type ItemStack struct {
	ID       uint      `json:"id" 			gorm:"primary_key"`
	Item     ItemModel `json:"item" 		gorm:"foreignkey:ID"`
	Number   uint      `json:"number"`
	Type     StackType `json:"type" 		gorm:"foreignkey:ID"`
	ItemID   uint      `json:"itemID" 		gorm:"column:item_id"`
	ChoiceID uint      `json:"choiceID" 	gorm:"column:choice_id"`
	TypeID   uint      `json:"typeID" 		gorm:"column:stack_type_id"`
}

func (stack ItemStack) ToDto() (ret struct {
	ID       uint      `json:"id"`
	Item     ItemModel `json:"item"`
	Number   uint      `json:"number"`
	Type     StackType `json:"type"`
	ItemID   uint      `json:"itemID"`
	ChoiceID uint      `json:"choiceID"`
	TypeID   uint      `json:"typeID"`
}) {
	ret.ID = stack.ID
	ret.Item = stack.Item
	ret.Number = stack.Number
	ret.Type = stack.Type
	ret.ItemID = stack.ItemID
	ret.ChoiceID = stack.ChoiceID
	ret.TypeID = stack.TypeID

	return
}
