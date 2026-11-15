//ff:func feature=token type=rule control=sequence
//ff:what TestKnownCandidatesUnit — a known word told that it begins right after a number with the letters of a unit and a word written after a unit (unit) keeps its candidate, its verdict lowered by that kind's rebuttal: 년간 of the small table for it stands below its verdict without the fact and above 0
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestKnownCandidatesUnit(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-unit-word.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	free := ix.knownCandidates("년간", rule.TailBare, false, false, false)
	after := ix.knownCandidates("년간", rule.TailBare, false, false, true)
	if len(free) != 1 || len(after) != 1 || free[0].name != "열쇠:체언" || after[0].name != free[0].name {
		t.Fatal(free, after)
	}
	if after[0].v <= 0 || after[0].v >= free[0].v/2 {
		t.Fatal(free[0].v, after[0].v)
	}
}
