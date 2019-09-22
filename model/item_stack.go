package model

// ItemStack is a pair of an item and a number of said items
type ItemStack struct {
	ID             uint      `json:"id" 		gorm:"primary_key"`
	ChoiceCostID   uint      `json:"choiceID" 	gorm:"column:choice_cost_id"`
	ChoiceRewardID uint      `json:"choiceID" 	gorm:"column:choice_reward_id"`
	ItemID         uint      `json:"itemID" 	gorm:"column:item_id"`
	Item           ItemModel `json:"item" 		gorm:"foreignkey:ID"`
	Number         uint      `json:"number"`
}

func (item ItemStack) ToDto() (ret struct {
	ID             uint      `json:"id"`
	ChoiceCostID   uint      `json:"choiceID"`
	ChoiceRewardID uint      `json:"choiceID"`
	ItemID         uint      `json:"itemID"`
	Item           ItemModel `json:"item"`
	Number         uint      `json:"number"`
}) {
	ret.ID = item.ID
	ret.ChoiceCostID = item.ChoiceCostID
	ret.ChoiceRewardID = item.ChoiceRewardID
	ret.ItemID = item.ItemID
	ret.Item = item.Item
	ret.Number = item.Number

	return
}
