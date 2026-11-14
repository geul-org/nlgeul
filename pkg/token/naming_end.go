//ff:func feature=token type=helper control=iteration dimension=1
//ff:what namingEnd — is r the last letter of a naming closing a particle chain is read right after (nounHeads): where such a chain is looked for; whether the closing before it is one of them is the graph's to judge
package token

import "strings"

func namingEnd(r rune) bool {
	for _, name := range nounHeads {
		if strings.HasSuffix(name, string(r)) {
			return true
		}
	}
	return false
}
