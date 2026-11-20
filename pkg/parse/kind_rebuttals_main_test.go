//ff:func feature=parse type=rule control=sequence
//ff:what TestKindRebuttalsMain — on a graph with the rebuttals of a kind, a one-syllable stem seen enough that is a main verb's beginning its unit with an auxiliary verb written onto its 어/아 form (Main) stands fully, where without the fact it stands lower; the fact does not lift "rarely seen", and a stem of more syllables seen enough stands fully either way
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestKindRebuttalsMain(t *testing.T) {
	enough := 10000
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	kindRebuttals(g, claim, enough)
	judge := func(form string, count int, main bool) float64 {
		v, _ := verdictOf(g, rule.Piece{Form: form, Count: count, Main: main}, "", false)
		return v
	}
	if short := judge("가", enough, false); short >= 1 || short <= 0 {
		t.Fatal("a one-syllable stem seen enough stands lower", short)
	}
	if v := judge("가", enough, true); v != 1 {
		t.Fatal("a main verb's stem beginning its unit before an auxiliary verb", v)
	}
	if few, main := judge("가르", 1, false), judge("가르", 1, true); main != few || few >= 1 {
		t.Fatal("seen once it is still rarely seen", few, main)
	}
	if v := judge("가르", enough, true); v != 1 {
		t.Fatal("a stem of more syllables seen enough", v)
	}
}
