//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWindowLongWord — a known word longer than the window is not cut by it wherever it stands in a long unit: with letters before it of every number from 0 to 14 (so that the window's end falls in every place of the word), the word is the letters of one node
package token

import (
	"strings"
	"testing"
)

func TestWindowLongWord(t *testing.T) {
	r := streamReader(t, "words-far.tsv")
	const word = "대한민국임시정부수립기념일"
	for k := 0; k < 15; k++ {
		text := windowText(2) + strings.Repeat("물", k) + word + "을마신다" + windowText(2)
		kept := false
		for _, line := range flatNodes(r, text, true) {
			kept = kept || strings.Contains(line, word+"[열쇠:체언]")
		}
		if !kept {
			t.Fatal(k, "the long word is cut")
		}
	}
}
