package arithmetic

type SingularArithmeticExpression struct {
	Value int
}

func (exp SingularArithmeticExpression) Calculate() int {
	return exp.Value
}
