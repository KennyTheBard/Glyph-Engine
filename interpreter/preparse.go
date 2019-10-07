package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

const queryIndentifier = '$'
const rollIdentifier = 'D'

func Preparse(rawScript string, context ExecutionContext) string {
	var scriptBuilder strings.Builder
	last := 0

	for i := 0; i < len(rawScript); i++ {
		if rawScript[i] == queryIndentifier {
			scriptBuilder.WriteString(rawScript[last:i])

			var q Query
			for j := i + 1; j < len(rawScript); j++ {
				if rawScript[j] == ']' {
					fmt.Sscanf(rawScript[i:j+1], "$%s.%s[\"%s\"]", &q.TableName, &q.StackType, &q.AttributeName)
					i = j
					break
				}
			}

			scriptBuilder.WriteString(strconv.Itoa(q.Execute(context)))

			last = i
		} else if rawScript[i] == rollIdentifier {
			scriptBuilder.WriteString(rawScript[last:i])

			for j := i + 1; j < len(rawScript); j++ {
				if !strings.ContainsRune("0123456789", rune(rawScript[j])) {
					if num, err := strconv.Atoi(rawScript[i+1 : j]); err == nil {
						scriptBuilder.WriteString(strconv.Itoa(Roll(uint(num))))
					}
					i = j
					break
				}
			}

			last = i
		}
	}

	scriptBuilder.WriteString(rawScript[last:])

	return scriptBuilder.String()
}
