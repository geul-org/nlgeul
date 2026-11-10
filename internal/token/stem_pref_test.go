//ff:func feature=token type=reader control=sequence
//ff:what TestStemPref — a known stem with a count: the log of its share; a known stem of words bound by a slash or a hyphen and of more than one token: the sum of its words'; a new word, a known one without a count, a bound stem that is one token (an atom), a bound stem not known, and a stem with both signs: its tokens
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStemPref(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if got := ix.stemPref("고양이", rule.TailCase, true, 3); math.Abs(got-math.Log10(1000/ix.lexTotal)) > 1e-9 {
		t.Fatal(got)
	}
	per := math.Log10(1.0 / syllables)
	share := func(w string) float64 { return math.Log10(float64(ix.lex[w].count) / ix.lexTotal) }
	if ix.stemPref("고양이", rule.TailCase, false, 3) != 3*per || ix.stemPref("것", rule.TailBare, true, 1) != per {
		t.Fatal(ix.stemPref("것", rule.TailBare, true, 1))
	}
	slash := share("입력") + per
	if got := ix.stemPref("입력/ab", rule.TailBare, true, 4); math.Abs(got-slash) > 1e-9 {
		t.Fatal("a stem of words bound by a slash", got, slash)
	}
	hyphen := share("입력") + share("출력")
	if got := ix.stemPref("입력-출력", rule.TailBare, true, 5); math.Abs(got-hyphen) > 1e-9 {
		t.Fatal("a stem of words bound by a hyphen", got, hyphen)
	}
	if got := ix.stemPref("ab/cd", rule.TailBare, true, 1); got != per {
		t.Fatal("a stem that is one token is not taken apart", got)
	}
	if got := ix.stemPref("입력/출력", rule.TailBare, true, 1); got != per {
		t.Fatal("one token even when its words are in the table", got)
	}
	if got := ix.stemPref("입력/ab", rule.TailBare, false, 4); got != 4*per {
		t.Fatal("a bound stem that is not known", got)
	}
	if got := ix.stemPref("입력/ab-cd", rule.TailBare, true, 6); got != 6*per {
		t.Fatal("a stem with both signs is not taken apart", got)
	}
	if got := ix.stemPref("입력출력", rule.TailBare, true, 4); got != 4*per {
		t.Fatal("a known stem with no count and no sign", got)
	}
}
