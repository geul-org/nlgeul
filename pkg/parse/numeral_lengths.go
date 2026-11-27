//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what numeralLengths — the lengths in bytes of the numerals of hangulNumerals that a text begins with; none when reached is false (nothing ends a run right before the text)
package parse

import "strings"

func numeralLengths(text string, reached bool) []int {
	var out []int
	for _, n := range hangulNumerals {
		if reached && strings.HasPrefix(text, n) {
			out = append(out, len(n))
		}
	}
	return out
}
