package data

import "errors"

type BaseEntity struct {
	ID uint `json:"id" gorm:"primary_key"`
}

func (entity *BaseEntity) Save() error {
	DB.Save(entity)
	return nil
}

func (entity *BaseEntity) FindById(id uint) error {
	if id == 0 {
		return errors.New("ID's must be positive numbers")
	}

	DB.First(entity, id)
	if entity.ID != id {
		return errors.New("No entity found with the given ID")
	}

	return nil
}

func (entity *BaseEntity) UpdateField(fieldName string, fieldValue interface{}) error {
	DB.Model(entity).Update(fieldName, fieldValue)

	return nil
}

func (entity *BaseEntity) UpdateFields(fields map[string]interface{}) error {
	for name, value := range fields {
		DB.Model(entity).Update(name, value)
	}

	return nil
}

func (entity *BaseEntity) Delete() error {
	DB.Delete(entity)
	return nil
}
