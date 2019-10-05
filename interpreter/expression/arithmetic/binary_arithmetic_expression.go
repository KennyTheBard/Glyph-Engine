package arithmetic

const (
	MULTIPLICATION = 0
	DIVISION       = 1
	MODULUS        = 2
)

type BinaryArithmeticExpression struct {
	Left           ArithmeticExpression
	Right          ArithmeticExpression
	ExpressionType uint
}

func (exp BinaryArithmeticExpression) Calculate() int {
	left := exp.Left.Calculate()
	right := exp.Right.Calculate()

	switch exp.ExpressionType {
	case MULTIPLICATION:
		return left * right
	case DIVISION:
		return left / right
	case MODULUS:
		return left % right
	default:
		return 0
	}
}
