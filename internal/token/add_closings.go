//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addClosings — the entries of every closing of the table; closings that need the copula stem are left out
package token

import "github.com/geul-org/koparse/internal/parse"

func (ix *index) addClosings() {
	for _, m := range parse.ClosingMarkers() {
		if m.Copula {
			continue
		}
		for _, e := range closingEntries(m) {
			ix.add(e)
		}
	}
}
