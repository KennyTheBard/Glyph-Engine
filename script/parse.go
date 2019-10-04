package script

import (
	exp "./expression"
)

type Branch struct {
	Cond         exp.Expression
	StoryID      uint
	ResolutionID uint
}
