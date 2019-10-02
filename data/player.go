package data

import "errors"

// UserModel is the main element of a page
type UserModel struct {
	BaseEntity
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserType    string
	CurrStoryID uint `json:"curr_story_id`
}

func (user UserModel) ToDto() (ret struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CurrStoryID uint   `json:"curr_story_id`
}) {
	ret.ID = user.ID
	ret.Username = user.Username
	ret.Password = user.Password
	ret.CurrStoryID = user.CurrStoryID

	return
}

func (user *UserModel) FindByUsername(username string) error {
	DB.Where("username = ?", username).First(user)
	if user.Username != username {
		return errors.New("No user found with the given ID")
	}

	return nil
}

func (user *UserModel) GetInventory() []ItemStack {
	var stacks []ItemStack
	DB.Where("owner_id = ? and type = ?", user.ID, OWNER_USER).Find(&stacks)
	return stacks
}

func (user *UserModel) GetCurrentStory() StoryModel {
	var story StoryModel
	story.FindById(user.CurrStoryID)
	return story
}
