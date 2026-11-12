//ff:func feature=token type=reader control=sequence
//ff:what TestStemOfTypoBound — the stem of a bound stem with a word written wrong is the right form of the whole when the tail asks so (boundTypo) and the letters come back as the fourth value; without that the stem is read as it was (no known word, 0); a stem written right has no written form
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStemOfTypoBound(t *testing.T) {
	sc := typoScan("입력/역활을", layerCost)
	eul := []piece{{"을", "목적격:을", 1, -1, rule.TailCase, false, ""}}
	stem, name, v, wrote := sc.stemOf(0, tail{start: 5, pieces: eul, boundTypo: true}, false)
	if stem != "입력/역할" || name != "대비:/" || v <= 0 || wrote != "입력/역활" {
		t.Fatal(stem, name, v, wrote)
	}
	stem, name, v, wrote = sc.stemOf(0, tail{start: 5, pieces: eul}, false)
	if stem != "입력/역활" || name != "" || v != 0 || wrote != "" {
		t.Fatal(stem, name, v, wrote)
	}
	right := typoScan("입력/출력을", layerCost)
	if stem, name, v, wrote = right.stemOf(0, tail{start: 5, pieces: eul, boundTypo: true}, false); stem != "입력/출력" || name != "대비:/" || v <= 0 || wrote != "" {
		t.Fatal(stem, name, v, wrote)
	}
	if stem, _, _, wrote = sc.stemOf(0, tail{start: 2, shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{tail: 8}}}, false); stem != "맡기" || wrote != "맞기" {
		t.Fatal(stem, wrote)
	}
}
