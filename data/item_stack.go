package data

const (
	OWNER_CHOICE = 1
	OWNER_USER   = 2
)

// ItemStack is a pair of an stack and a number of said items
type ItemStack struct {
	BaseEntity
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
