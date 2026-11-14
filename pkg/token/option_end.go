//ff:func feature=token type=helper control=iteration dimension=1
//ff:what optionEnd — where a command's option ends when a unit begins with one: one or two dashes, a Latin letter, then Latin letters, digits and - _ =; 0 when the unit begins with none
package token

import (
	"strings"
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func optionEnd(s []rune) int {
	i := len(s) - len([]rune(strings.TrimLeft(string(s), "-")))
	if i < 1 || i > 2 || i >= len(s) || s[i] > unicode.MaxASCII || !unicode.IsLetter(s[i]) {
		return 0
	}
	for i < len(s) && s[i] < unicode.MaxASCII && (rule.RunRune(s[i]) || strings.ContainsRune("-=", s[i])) {
		i++
	}
	return i
}
