//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what syllablesFrom — how many Hangul syllables stand in a row from a place of a text
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func syllablesFrom(s []rune, from int) int {
	n := 0
	for from+n < len(s) && rule.JongOf(s[from+n]) >= 0 {
		n++
	}
	return n
}
