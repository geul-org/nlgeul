//ff:func feature=token type=helper control=iteration dimension=1
//ff:what waitsFor — does some rule of a sentence that speaks of the candidate need the words after it, up to the predicate, to judge it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func waitsFor(self rule.Near, rules []sentenceRule) bool {
	for _, r := range rules {
		if r.waits != nil && r.about(self) && r.waits(self) {
			return true
		}
	}
	return false
}
