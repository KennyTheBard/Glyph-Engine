package arithmetic

type ValueArithmeticExpression struct {
	Value int
}

func (exp ValueArithmeticExpression) Calculate() int {
	return exp.Value
}
