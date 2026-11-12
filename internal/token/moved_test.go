//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestMoved — of the candidates of a unit headed by a sign, those whose first word starts with the sign are carried over; a candidate that took the sign into a word's letters, and one with no word, are left
package token

import "testing"

func TestMoved(t *testing.T) {
	ix, _ := newIndex(Options{})
	sc := unitScan(ix, unit{text: ")를", head: 1, name: pairName, reads: []pairing{{shutName, 1}}})
	got := sc.moved(sc.search(), 1)
	if len(got) == 0 {
		t.Fatal("none carried over")
	}
	for _, p := range got {
		w := wordsOf(p)
		if !w[0].word.Attach || w[0].word.Parts[0].Text != "를" || w[0].end != 1 {
			t.Fatal(w)
		}
	}
	if sc.moved(nil, 1) != nil || sc.moved([]*path{{verdict: 1}}, 1) != nil {
		t.Fatal("no candidate, no word")
	}
}
