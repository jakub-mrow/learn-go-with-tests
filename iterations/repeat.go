package iteration

import "strings"

func Repeat(character string, repeateTimes int) string {
	var repeated strings.Builder
	for i := 0; i < repeateTimes; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
