package validator

import "reflect"

type Validator struct {
	ValidationType reflect.Type

	// ValidationFunc should make an internal type assertion and return nil if it fails,
	// otherwise proceed with validation checks
	ValidationFunc func(object interface{}, validationType reflect.Type) error
}

var Validators []Validator

func Validate(object interface{}) error {
	var err error
	for _, validator := range Validators {
		if err = validator.ValidationFunc(object, validator.ValidationType); err != nil {
			return err
		}
	}

	return nil
}
