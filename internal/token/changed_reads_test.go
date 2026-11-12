//ff:func feature=token type=reader control=sequence
//ff:what TestChangedReads — a tail is read once more for the stem of the word table that the letters are a changed form of (들 before 었 is 듣), keeping the reads already there; none when the tail is not one the stem takes
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestChangedReads(t *testing.T) {
	sc := wordsScan("들었다", layerCost)
	past := []piece{{text: "었", tail: rule.TailPred}, {text: "다", tail: rule.TailPred}}
	got := sc.changedReads([]tail{{start: 1}}, "들", tail{start: 1, pieces: past, head: headWhole})
	if len(got) != 2 || got[0].shown != "" || got[1].shown != "듣" || got[1].irr != 0 {
		t.Fatal(got)
	}
	if got = sc.changedReads(nil, "들", tail{start: 1, pieces: []piece{{text: "이", tail: rule.TailCase}}, head: headWhole}); len(got) != 0 {
		t.Fatal(got)
	}
}
