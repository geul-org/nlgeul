//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestSuffixGateOf — the graph is named as the nominal gate; it passes a nominal of the word table whose stem and suffix together are no word of the table, and turns down a word that is no nominal or whose stem and suffix together are a word, each with its reason
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSuffixGateOf(t *testing.T) {
	e := suffixGateOf()
	if e.name != nominalGate || e.g == nil {
		t.Fatal(e)
	}
	for _, c := range []struct {
		p    rule.Piece
		pass bool
		why  string
	}{
		{rule.Piece{Form: "보장", HasStem: true, Next: -1, Stem: kindNominal}, true, ""},
		{rule.Piece{Form: "보장", HasStem: true, Next: -1, Stem: kindNominal | kindVerb}, true, ""},
		{rule.Piece{Form: "합격", HasStem: true, Next: -1, Stem: kindNominal, Whole: true}, false, whyJoinedWord},
		{rule.Piece{Form: "가", HasStem: true, Next: -1, Stem: kindVerb}, false, whyKind},
		{rule.Piece{Form: "뷁", HasStem: true, Next: -1}, false, whyKind},
	} {
		v, why := verdictOf(e.g, c.p, "", true)
		if (v > 0) != c.pass || (!c.pass && !slices.Contains(why, c.why)) {
			t.Errorf("%+v: %v %v", c.p, v, why)
		}
	}
}
