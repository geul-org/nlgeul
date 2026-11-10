//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddClosings — the closings are indexed by their last token or their lead's key; those that need the copula stem are not there
package token

import "testing"

func TestAddClosings(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	if len(ix.byKey['만']) == 0 || len(ix.byKey['ㄴ']) == 0 || len(ix.byKey[keyVowel]) == 0 {
		t.Fatal("missing keys")
	}
	for _, e := range ix.byKey['요'] {
		if e.form == "에요" || e.form == "예요" {
			t.Fatal(e.name)
		}
	}
}
