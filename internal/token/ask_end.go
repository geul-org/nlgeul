//ff:func feature=token type=helper control=iteration dimension=1
//ff:what askEnd — is r the last letter of a closing that asks, which a particle is read right after (askClosings): where such a particle is looked for; whether the closing before it is one of them is the graph's to judge
package token

import "strings"

func askEnd(r rune) bool {
	for _, name := range askClosings {
		if strings.HasSuffix(name, string(r)) {
			return true
		}
	}
	return false
}
