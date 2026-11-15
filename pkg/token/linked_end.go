//ff:func feature=token type=helper control=iteration dimension=1
//ff:what linkedEnd — is r the last letter of a closing the copula stands right after (copulaClosings): where the closing before a copula stem is looked for; whether the tail before it ends with one of them is afterLinked's to tell
package token

import "strings"

func linkedEnd(r rune) bool {
	for _, name := range copulaClosings {
		if strings.HasSuffix(name, string(r)) {
			return true
		}
	}
	return false
}
