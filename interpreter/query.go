package interpreter

import (
	"strings"

	data "../data"
)

type ExecutionContext struct {
	User   data.UserModel
	Choice data.ChoiceModel
}

type Query struct {
	TableName     string
	StackType     string
	AttributeName string
}

func (q Query) Execute(context ExecutionContext) int {
	var stacks []data.AttributeStack
	switch strings.ToLower(q.TableName) {
	case "user":
		stacks = context.User.GetInventory()

	case "choice":
		stacks = context.Choice.GetAttributeStacks()

	default:
		stacks = []data.AttributeStack{}
	}

	for _, stack := range stacks {
		if stack.StackType == q.StackType {
			attr, err := stack.GetAttribute()
			if err != nil {
				break
			}

			if attr.Name == q.AttributeName {
				return int(stack.Number)
			}
		}
	}

	return 0 // treat error
}
