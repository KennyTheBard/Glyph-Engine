package expression

const (
	OR  = 0
	AND = 1
)

type LogicExpression struct {
	Subs           []Expression
	ExpressionType uint
}

func (exp LogicExpression) Evaluate() bool {
	for _, sub := range exp.Subs {
		value := sub.Evaluate()

		if exp.ExpressionType == OR && value {
			return true
		}

		if exp.ExpressionType == AND && !value {
			return false
		}
	}

	return true
}
