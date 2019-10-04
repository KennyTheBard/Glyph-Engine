package arithmetic

const (
	ADDITION     = 0
	SUBSTRACTION = 1
)

type MultiArithmeticExpression struct {
	Subs           []ArithmeticExpression
	ExpressionType uint
}

func (exp MultiArithmeticExpression) Calculate() int {
	total := 0

	for _, sub := range exp.Subs {
		value := sub.Calculate()

		if exp.ExpressionType == ADDITION {
			total += value
		} else if exp.ExpressionType == SUBSTRACTION {
			total -= value
		}
	}

	return total
}
