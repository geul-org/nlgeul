//ff:func feature=parse type=rule control=sequence
//ff:what TestAdverbExcept — four edges on the graph of an adverb's stem: the copula after it stands, lower than the adverb with an auxiliary particle (the lift's reason in the trace); it does not stand when the tables give the stem as a nominal too, nor when the stem is a run of Latin letters, each with its reason; a case particle after it does not stand
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAdverbExcept(t *testing.T) {
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	tail := except(g, rule.StemTail, whyStemTail, &rule.Bits{Of: "tail", Mask: kindTails[kindAdverb]})
	tail.Attacks(claim)
	if n := adverbExcept(g, claim, tail); n != 4 {
		t.Fatal("edges", n)
	}
	aux, _ := verdictOf(g, rule.Piece{Form: "아직", Tail: rule.TailAux, Stem: kindAdverb}, "", true)
	copula, why := verdictOf(g, rule.Piece{Form: "아직", Tail: rule.TailCopula, Stem: kindAdverb}, "", true)
	if aux != 1 || copula <= 0 || copula >= aux || !slices.Contains(why, whyAdverbLoose) {
		t.Fatal("아직 before the copula", aux, copula, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "물론", Tail: rule.TailCopula, Stem: kindAdverb | kindNominal}, "", true); v > 0 || !slices.Contains(why, whyAlsoNominal) {
		t.Fatal("a nominal too", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "off", Tail: rule.TailCopula, Stem: kindAdverb}, "", true); v > 0 || !slices.Contains(why, whyStemRun) {
		t.Fatal("a run of Latin letters", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "아직", Tail: rule.TailCase, Stem: kindAdverb}, "", true); v > 0 || !slices.Contains(why, whyStemTail) {
		t.Fatal("a case particle", v, why)
	}
}
