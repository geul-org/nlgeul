//ff:func feature=token type=reader control=sequence
//ff:what TestPartsPref — the sum of the words' preferences: a word of the table by its count (a run of Latin letters in the table too), another by its tokens: a run of Latin letters and digits is one, a word of Hangul or of Hangul and Latin one for each letter; no words, nothing
package token

import (
	"math"
	"testing"
)

func TestPartsPref(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	per := math.Log10(1.0 / syllables)
	share := func(w string) float64 { return math.Log10(float64(ix.lex[w].count) / ix.lexTotal) }
	want := share("입력") + per
	if got := ix.partsPref([]string{"입력", "ab"}); math.Abs(got-want) > 1e-9 {
		t.Fatal(got, want)
	}
	if got := ix.partsPref([]string{"입력", "출력"}); math.Abs(got-(share("입력")+share("출력"))) > 1e-9 {
		t.Fatal("two words of the table", got)
	}
	if got := ix.partsPref([]string{"사자", "ab"}); math.Abs(got-3*per) > 1e-9 {
		t.Fatal("a word with no count costs its letters, a run one token", got)
	}
	if got := ix.partsPref([]string{"ab_1", "x2"}); math.Abs(got-2*per) > 1e-9 {
		t.Fatal("runs of Latin letters, digits and underscores", got)
	}
	if got := ix.partsPref([]string{"곰a", "입력"}); math.Abs(got-(2*per+share("입력"))) > 1e-9 {
		t.Fatal("a word of Hangul and Latin letters costs each letter", got)
	}
	if got := ix.partsPref([]string{"foo", "ab"}); math.Abs(got-(share("foo")+per)) > 1e-9 {
		t.Fatal("a run the table lists goes by its count", got)
	}
	if got := ix.partsPref(nil); got != 0 {
		t.Fatal("no words", got)
	}
}
