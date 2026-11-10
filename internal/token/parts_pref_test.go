//ff:func feature=token type=reader control=sequence
//ff:what TestPartsPref — the sum of the words' preferences: a word of the table by its count, another by its tokens (a run of Latin letters is one); a stem that is one atom is not taken apart
package token

import (
	"math"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestPartsPref(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	want := math.Log10(6000/ix.lexTotal) + math.Log10(1.0/syllables)
	if got := ix.partsPref([]string{"입력", "ab"}); math.Abs(got-want) > 1e-9 {
		t.Fatal(got, want)
	}
	if got := ix.stemPref("입력/ab", rule.TailBare, true, 4); math.Abs(got-want) > 1e-9 {
		t.Fatal("a stem of bound words", got)
	}
	if got := ix.partsPref([]string{"사자", "ab"}); math.Abs(got-3*math.Log10(1.0/syllables)) > 1e-9 {
		t.Fatal("a word with no count costs its letters, a run one token", got)
	}
	if got := ix.stemPref("ab/cd", rule.TailBare, true, 1); math.Abs(got-math.Log10(1.0/syllables)) > 1e-9 {
		t.Fatal("an atom is one token and is not taken apart", got)
	}
}
