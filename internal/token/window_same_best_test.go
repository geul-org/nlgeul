//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWindowSameBest — the best way is the same with the window on and off: the nodes of the best ways of the stretches, from the first to the last, are the same words at the same places for made-up long units (words written together, with signs between letters, a sign of a pair with letters written onto it, a long known word at every place of the window)
package token

import (
	"strings"
	"testing"
)

func TestWindowSameBest(t *testing.T) {
	r := streamReader(t, "words-far.tsv")
	texts := []string{windowText(5), strings.Repeat("고양이/마당에서물·사과를마신다", 4), "(고양이)" + windowText(4), windowText(2) + "대한민국임시정부수립기념일을마신다" + windowText(2), "고양이가" + windowText(3)}
	for _, text := range texts {
		off, on := flatNodes(r, text, false), flatNodes(r, text, true)
		if len(on) == 0 || strings.Join(on, "\n") != strings.Join(off, "\n") {
			t.Fatal(text, "\n", on, "\n", off)
		}
	}
}
