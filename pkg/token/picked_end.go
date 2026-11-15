//ff:func feature=token type=helper control=iteration dimension=1
//ff:what pickedEnd — is r the last letter of a closing the particle of this form is written onto (pickedClosings): where such a particle is looked for; whether the closing before it is one of them is the graph's to judge
package token

import "strings"

func pickedEnd(form string, r rune) bool {
	for _, name := range pickedClosings[form] {
		if strings.HasSuffix(name, string(r)) {
			return true
		}
	}
	return false
}
