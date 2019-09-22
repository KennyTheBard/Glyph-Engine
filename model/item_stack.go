package model

// ItemStack is a pair of an item and a number of said items
type ItemStack struct {
	ID       uint      `json:"id" 			gorm:"primary_key"`
	ItemID   uint      `json:"itemID" 		gorm:"column:item_id"`
	Item     ItemModel `json:"item" 		gorm:"foreignkey:ID"`
	Number   uint      `json:"number"`
	ChoiceID uint      `json:"choiceID" 	gorm:"column:choice_id"`
}

func (stack ItemStack) ToDto() (ret struct {
	ID       uint      `json:"id"`
	ItemID   uint      `json:"itemID"`
	Item     ItemModel `json:"item"`
	Number   uint      `json:"number"`
	ChoiceID uint      `json:"choiceID"`
}) {
	ret.ID = stack.ID
	ret.ItemID = stack.ItemID
	ret.Item = stack.Item
	ret.Number = stack.Number
	ret.ChoiceID = stack.ChoiceID

	return
}
