//ff:func feature=token type=helper control=sequence
//ff:what partTokens — how many tokens one word of a bound stem holds for preference: a run of Latin letters and digits is an atom and counts as one; else one for each letter
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func partTokens(part string) int {
	if strings.IndexFunc(part, func(r rune) bool { return !rule.RunRune(r) }) < 0 {
		return 1
	}
	return len([]rune(part))
}
