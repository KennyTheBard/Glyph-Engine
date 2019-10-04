package arithmetic

type QueryArithmeticExpression struct {
	QueryFunc func(string) int
	Query     string
}

func (exp QueryArithmeticExpression) Calculate() int {
	return exp.QueryFunc(exp.Query)
}
