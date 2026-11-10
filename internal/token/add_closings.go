//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addClosings — the entries of every closing of the table; the names of those the copula does not take are kept for its graph
package token

import "github.com/geul-org/koparse/internal/parse"

func (ix *index) addClosings() {
	for _, m := range parse.ClosingMarkers() {
		for _, e := range closingEntries(m) {
			ix.add(e)
			if e.bars {
				ix.barred = append(ix.barred, e.name)
			}
		}
	}
}
