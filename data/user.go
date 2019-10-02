package data

import "errors"

// UserModel is the main element of a page
type UserModel struct {
	ID          uint   `json:"id" gorm:"primary_key"`
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
	DB.Where("owner_id = ? and owner_type = ?", user.ID, OWNER_USER).Find(&stacks)
	return stacks
}

func (user *UserModel) GetCurrentStory() (StoryModel, error) {
	var story StoryModel
	err := story.FindById(user.CurrStoryID)
	return story, err
}

// CRUD methods

func (user *UserModel) Save() error {
	DB.Save(user)
	return nil
}

func (user *UserModel) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(user, id)
	if user.ID != id {
		return errors.New("No user found with the given ID")
	}

	return nil
}

func (user *UserModel) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(user).Update(fieldName, fieldValue)

	return nil
}

func (user *UserModel) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(user).Update(name, value)
	}

	return nil
}

func (user *UserModel) Delete() error {
	DB.Delete(user)
	return nil
}
