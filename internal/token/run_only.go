//ff:func feature=token type=helper control=iteration dimension=1
//ff:what runOnly — whether a text is made of the letters and digits of a run of Latin letters and digits alone (a part of a bound stem that is such a run stands as one, partVerdict)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func runOnly(text string) bool {
	for _, r := range text {
		if !rule.RunRune(r) {
			return false
		}
	}
	return true
}
