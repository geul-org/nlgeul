//ff:func feature=token type=helper control=sequence
//ff:what closesOf — for a piece of a tail that is its closing (named 맺음): whether the closing ends a sentence, links, or both, as the closing tables mark it (rule.KindFinal, rule.KindLink); 0 for any other piece and for a tail with no closing entry
package token

import "strings"

func closesOf(t tail, p piece) int {
	if t.close == nil || !strings.HasPrefix(p.names, closingPrefix) {
		return 0
	}
	return t.close.ends
}
