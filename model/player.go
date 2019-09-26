package model

// PlayerModel is the main element of a page
type PlayerModel struct {
	ID       uint        `json:"id" 		gorm:"primary_key"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Items    []ItemStack `json:"items"`
}

func (player PlayerModel) ToDto() (ret struct {
	ID       uint        `json:"id" 		gorm:"primary_key"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Items    []ItemStack `json:"items"`
}) {
	ret.ID = player.ID
	ret.Username = player.Username
	ret.Password = player.Password
	ret.Items = player.Items

	return
}
