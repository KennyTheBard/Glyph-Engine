package model

import (
	"github.com/jinzhu/gorm"
)

// Choice is the main subelement of the page
type Requirement struct {
	gorm.Model

	// id of the choice record
	Choice uint

	// id of the item record
	Item uint

	// minimum number of items needed
	Number uint

	// if true => Number is the maximum number of items needed
	AtMost bool
}
