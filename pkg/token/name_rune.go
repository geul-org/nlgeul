//ff:func feature=token type=helper control=sequence
//ff:what nameRune — is a letter one a path's name is written with besides Hangul: a Latin letter, a digit, or one of . ~ _ -
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func nameRune(r rune) bool {
	return rule.RunRune(r) || strings.ContainsRune(".~_-", r)
}
