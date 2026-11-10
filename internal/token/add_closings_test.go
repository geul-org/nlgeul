//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddClosings — the closings are indexed by their last token or their lead's key; those that attach only to the copula stem are there too, and the names of those it does not take are kept
package token

import (
	"slices"
	"testing"
)

func TestAddClosings(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	if len(ix.byKey['만']) == 0 || len(ix.byKey['ㄴ']) == 0 || len(ix.byKey[keyVowel]) == 0 {
		t.Fatal("missing keys")
	}
	var needs, holds bool
	for _, e := range ix.byKey['요'] {
		needs, holds = needs || (e.form == "에요" && e.needs), holds || (e.form == "예요" && e.holds)
	}
	if !needs || !holds || !slices.Contains(ix.barred, "꾸밈:는") || slices.Contains(ix.barred, "맺음:다") {
		t.Fatal(needs, holds, ix.barred)
	}
}
