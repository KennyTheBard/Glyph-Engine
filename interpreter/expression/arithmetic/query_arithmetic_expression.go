package arithmetic

type QueryArithmeticExpression struct {
	QueryFunc func(string) int
	Statement string
}

func (exp QueryArithmeticExpression) Calculate() int {
	return exp.QueryFunc(exp.Statement)
}
