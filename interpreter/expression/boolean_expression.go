package expression

import (
	arith "./arithmetic"
)

const (
	EQUAL                 = 0
	NOT_EQUAL             = 1
	GREATER_THAN          = 2
	GREATER_OR_EQUAL_THAN = 3
	LESSER_THAN           = 4
	LESSER_OR_EQUAL_THAN  = 5
)

type BooleanExpression struct {
	Left           arith.ArithmeticExpression
	Right          arith.ArithmeticExpression
	ExpressionType int
}

func (exp BooleanExpression) Evaluate() bool {
	left := exp.Left.Calculate()
	right := exp.Right.Calculate()

	switch exp.ExpressionType {
	case EQUAL:
		return left == right
	case NOT_EQUAL:
		return left != right
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
