//ff:func feature=token type=helper control=iteration dimension=1
//ff:what isPath — is a text a path: it begins with a path's head (pathHeads: /, ~/, ./, ../), or it has two slashes or more with a name's letter (Latin letter, digit, . ~ _ -) on both sides and is not numbers and slashes only
package token

import (
	"slices"
	"strings"
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func isPath(s []rune) bool {
	text := string(s)
	if slices.ContainsFunc(pathHeads, func(head string) bool { return strings.HasPrefix(text, head) }) {
		return true
	}
	slashes, letters := 0, false
	for i, r := range s {
		letters = letters || unicode.IsLetter(r)
		if r == '/' && i > 0 && i+1 < len(s) && nameRune(s[i-1]) && nameRune(s[i+1]) {
			slashes++
		}
	}
	return slashes >= 2 && letters && rule.JongOf(s[0]) < 0
}
