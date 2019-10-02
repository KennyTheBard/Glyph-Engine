package data

// Item is the principal form of resource in the game
type ItemModel struct {
	BaseEntity
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type`
}

func (item ItemModel) ToDto() (ret struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type`
}) {
	ret.ID = item.ID
	ret.Name = item.Name
	ret.Text = item.Text
	ret.Type = item.Type

	return
}
