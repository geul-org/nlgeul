//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lastOfAny — is r the last letter of one of some closings (by their names: askClosings, copulaClosings, pickedClosings of a form, chainHeads): where what is written right after such a closing is looked for; whether the closing before it is one of them is the graph's to judge
package token

import "strings"

func lastOfAny(names []string, r rune) bool {
	for _, name := range names {
		if strings.HasSuffix(name, string(r)) {
			return true
		}
	}
	return false
}
