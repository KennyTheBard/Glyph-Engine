package interpreter

import (
	"strconv"
	"strings"

	exp "./expression"
	arith "./expression/arithmetic"
)

type Branch struct {
	condition       exp.Expression
	storyID         uint
	hasResolutionID bool
	resolutionID    uint
}

func Parse(rawText string) (Script, error) {
	var script Script
	scriptText := strings.Join(strings.Fields(rawText), "")

	lines := strings.FieldsFunc(scriptText, ContainsRunes(LINE_DELIMITATORS))
	script.Branches = make([]Branch, len(lines))

	for i, line := range lines {
		lineTokens := strings.FieldsFunc(line, ContainsRunes(INLINE_DELIMITATORS))

		script.Branches[i].condition = parseExpression(lineTokens[0])

		if id, err := strconv.Atoi(lineTokens[1]); err == nil {
			script.Branches[i].storyID = uint(id)
		} else {
			return script, err
		}

		if len(lineTokens) == 3 {
			script.Branches[i].hasResolutionID = true
			if id, err := strconv.Atoi(lineTokens[2]); err == nil {
				script.Branches[i].resolutionID = uint(id)
			} else {
				return script, err
			}
		}
	}

	return script, nil
}

func parseExpression(expText string) exp.Expression {
	isOr := strings.Index(expText, LOGIC_OR_DELIM) != -1
	isAnd := strings.Index(expText, LOGIC_AND_DELIM) != -1

	if isOr && isAnd {
		return exp.LogicExpression{[]exp.Expression{}, exp.AND_OR}
	}

	if isOr || isAnd {
		var expression exp.LogicExpression
		var subexps []string

		if isOr {
			subexps = strings.Split(expText, LOGIC_OR_DELIM)
			expression.ExpressionType = exp.OR

		} else if isAnd {
			subexps = strings.Split(expText, LOGIC_AND_DELIM)
			expression.ExpressionType = exp.AND
		}

		expression.Subs = make([]exp.Expression, len(subexps))
		for i, subexp := range subexps {
			expression.Subs[i] = parseExpression(subexp)
		}

		return expression

	} else {
		var expression exp.BooleanExpression

		indexes := FindStrings(expText, BOOL_DELIMITATORS)
		if length := len(indexes); length > 1 {
			return expression // treat error
		} else if length < 1 {
			return expression // treat error
		}

		expression.Left = parseArithmeticExpression(expText[:indexes[0]])
		expression.Right = parseArithmeticExpression(expText[indexes[0]+2:])
		for i, delim := range BOOL_DELIMITATORS {
			if delim == expText[indexes[0]:indexes[0]+2] {
				expression.ExpressionType = i
			}
		}

		return expression
	}
}

func parseArithmeticExpression(expTest string) arith.ArithmeticExpression {

}
