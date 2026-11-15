//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestJoinedGateOf — the graph is named as asked; with no kinds asked it passes any stem, a new word too, and turns down one whose stem and piece together are a word of the table, with that reason; with kinds asked it also turns down a stem of the table with none of those kinds, with its own reason
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestJoinedGateOf(t *testing.T) {
	bare, nominal := joinedGateOf(joinedGate, 0), joinedGateOf(nominalGate, kindNominal)
	if bare.name != joinedGate || bare.g == nil || nominal.name != nominalGate || nominal.g == nil {
		t.Fatal(bare, nominal)
	}
	for _, c := range []struct {
		e    *entry
		p    rule.Piece
		pass bool
		why  string
	}{
		{bare, rule.Piece{Form: "학교", HasStem: true, Next: -1, Stem: kindNominal}, true, ""},
		{bare, rule.Piece{Form: "뷁", HasStem: true, Next: -1}, true, ""},
		{bare, rule.Piece{Form: "요구", HasStem: true, Next: -1, Stem: kindNominal, Whole: true}, false, whyJoinedParticle},
		{nominal, rule.Piece{Form: "여기", HasStem: true, Next: -1, Stem: kindNominal}, true, ""},
		{nominal, rule.Piece{Form: "나이", HasStem: true, Next: -1, Stem: kindNominal, Whole: true}, false, whyJoinedParticle},
		{nominal, rule.Piece{Form: "가", HasStem: true, Next: -1, Stem: kindVerb}, false, whyKind},
		{nominal, rule.Piece{Form: "뷁", HasStem: true, Next: -1}, false, whyKind},
	} {
		v, why := verdictOf(c.e.g, c.p, "", true)
		if (v > 0) != c.pass || (!c.pass && !slices.Contains(why, c.why)) {
			t.Errorf("%s %+v: %v %v", c.e.name, c.p, v, why)
		}
	}
}
