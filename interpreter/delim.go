package interpreter

import (
	"regexp"
	"strings"
)

// general purpose delimitors
var LINE_DELIMITATORS []rune = []rune{';'}
var INLINE_DELIMITATORS []rune = []rune{'?', ':'}

// logic expressions delimitators
const LOGIC_OR_DELIM = "||"
const LOGIC_AND_DELIM = "&&"

// boolean expressions delimitators
var BOOL_DELIMITATORS []string = []string{"==", "!=", ">>", ">=", "<<", "<="}

// multi operator arithmetic operations
var MULTI_ARITHMETIC_DELIMITERS []rune = []rune{'+', '-'}

func ContainsRunes(delimiters []rune) func(rune) bool {
	return func(r rune) bool {
		for _, delim := range delimiters {
			if r == delim {
				return true
			}
		}
		return false
	}
}

// FindStrings returns the starting index of any matched string fromm delimiters in the given sentence
func FindStrings(sentence string, delimiters []string) []int {
	matches := regexp.MustCompile(strings.Join(delimiters, "|")).FindAllStringIndex(sentence, -1)
	if matches == nil {
		return nil
	} else {
		return matches[:][0]
	}
}
