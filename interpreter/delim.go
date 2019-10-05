package interpreter

// general purpose delimitors
var LINE_DELIMITATORS []rune = []rune{';'}
var INLINE_DELIMITATORS []rune = []rune{'?', ':'}

// logic expressions delimitators
const LOGIC_OR_DELIM = "||"
const LOGIC_AND_DELIM = "&&"

// boolean expressions delimitators
var BOOL_DELIMITATORS []string = []string{"==", "!=", ">>", ">=", "<<", "<="}

const BOOL_EQ_IDX = 0
const BOOL_NE_IDX = 1
const BOOL_GT_IDX = 2
const BOOL_GE_IDX = 3
const BOOL_LT_IDX = 4
const BOOL_LE_IDX = 5

func IsRuneIn(delimiters []rune) func(rune) bool {
	return func(r rune) bool {
		for _, delim := range delimiters {
			if r == delim {
				return true
			}
		}
		return false
	}
}

func FindRune