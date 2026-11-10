//ff:func feature=token type=reader control=sequence
//ff:what TestCountOf — a kind a key has rows of has their count (0 when the rows give none); a kind it has none of has -1
package token

import "testing"

func TestCountOf(t *testing.T) {
	w := folded(folded(word{}, kindNominal, 7), kindAdverb, 0)
	if countOf(w, kindNominal, 0) != 7 || countOf(w, kindAdverb, 3) != 0 || countOf(w, kindVerb, 1) != -1 {
		t.Fatal(w)
	}
}
