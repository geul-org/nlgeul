//ff:func feature=token type=reader control=sequence
//ff:what TestFolded — a line adds its kind, keeps the greatest count and sums the count of its own kind
package token

import "testing"

func TestFolded(t *testing.T) {
	w := folded(word{}, kindVerb, irrD, 900)
	w = folded(w, kindNominal, 0, 40)
	w = folded(w, kindVerb, 0, 100)
	if w.kinds != kindVerb|kindNominal || w.count != 900 || w.by[0] != 40 || w.by[1] != 1000 || w.irr != irrD {
		t.Fatalf("%+v", w)
	}
}
