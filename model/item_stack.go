package model

// ItemStack is a pair of an item and a number of said items
type ItemStack struct {
	ID             uint      `json:"id" 		gorm:"primary_key"`
	ChoiceCostID   uint      `json:"choiceID" gorm:"column:choice_cost_id"`
	ChoiceRewardID uint      `json:"choiceID" gorm:"column:choice_reward_id"`
	ItemID         uint      `json:"itemID" 	gorm:"column:item_id"`
	Item           ItemModel `json:"item" 	gorm:"foreignkey:ID"`
	Num            uint      `json:"num"`
}

func (item ItemStack) ToDto() (ret struct {
	ID   uint      `json:"id"`
	Item ItemModel `json:"item"`
	Num  uint      `json:"num"`
}) {
	ret.ID = item.ID
	ret.Item = item.Item
	ret.Num = item.Num

	return
}
