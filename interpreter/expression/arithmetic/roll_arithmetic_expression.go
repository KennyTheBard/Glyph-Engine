package arithmetic

import (
	rand "math/rand"
)

type RollArithmeticExpression struct {
	Gen *rand.Rand
	Max int
}

func (exp RollArithmeticExpression) Calculate() int {
	if exp.Max <= 0 {
		return 0
	}

	return exp.Gen.Intn(exp.Max) + 1
}
