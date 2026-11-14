//ff:func feature=token type=helper control=sequence
//ff:what TestBornOf — the number of birth kept with a branch of one word; noBorn for one made without scores and for one with no word
package token

import "testing"

func TestBornOf(t *testing.T) {
	kept := Branch{Words: make([]Eojeol, 1), Verdict: 1, scores: []wordScore{{1, -1, 7}}}
	if bornOf(kept) != 7 || bornOf(Branch{Words: make([]Eojeol, 1), Verdict: 1}) != noBorn || bornOf(Branch{}) != noBorn {
		t.Fatal(bornOf(kept))
	}
}
