//ff:func feature=token type=reader control=sequence
//ff:what TestStemPref — a known stem with a count: the log of its share; a new word, or a known one without a count, its tokens
package token

import (
	"math"
	"testing"
)

func TestStemPref(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if got := ix.stemPref("고양이", true, 3); math.Abs(got-math.Log10(1000/ix.lexTotal)) > 1e-9 {
		t.Fatal(got)
	}
	per := math.Log10(1.0 / syllables)
	if ix.stemPref("고양이", false, 3) != 3*per || ix.stemPref("것", true, 1) != per {
		t.Fatal(ix.stemPref("것", true, 1))
	}
}
