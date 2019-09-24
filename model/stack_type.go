package model

// StackType holds the possible purposes an item stack may have
type StackType struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (stackType StackType) ToDto() (ret struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}) {
	ret.ID = stackType.ID
	ret.Name = stackType.Name

	return
}
