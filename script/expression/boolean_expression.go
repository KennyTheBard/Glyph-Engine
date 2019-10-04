package expression

import (
	arith "./arithmetic"
)

const (
	EQUAL                 = 0
	GREATER_THAN          = 1
	GREATER_OR_EQUAL_THAN = 2
	LESSER_THAN           = 3
	LESSER_OR_EQUAL_THAN  = 4
)

type BooleanExpression struct {
	Left           arith.ArithmeticExpression
	Right          arith.ArithmeticExpression
	ExpressionType uint
}

func (exp BooleanExpression) Evaluate() bool {
	left := exp.Left.Calculate()
	right := exp.Right.Calculate()

	switch exp.ExpressionType {
	case EQUAL:
		return left == right
	case GREATER_THAN:
		return left > right
	case GREATER_OR_EQUAL_THAN:
		return left >= right
	case LESSER_THAN:
		return left < right
	case LESSER_OR_EQUAL_THAN:
		return left <= right
	default:
		return true
	}
}
