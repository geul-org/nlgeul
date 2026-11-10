//ff:func feature=token type=reader control=sequence
//ff:what TestPartsPref — the sum of the words' preferences: a word of the table by its count, another by its letters
package token

import (
	"math"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestPartsPref(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	want := math.Log10(6000/ix.lexTotal) + 2*math.Log10(1.0/syllables)
	if got := ix.partsPref([]string{"입력", "ab"}); math.Abs(got-want) > 1e-9 {
		t.Fatal(got, want)
	}
	if got := ix.stemPref("입력/ab", rule.TailBare, true, 5); math.Abs(got-want) > 1e-9 {
		t.Fatal("a stem of bound words", got)
	}
}
