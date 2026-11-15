//ff:func feature=token type=rule control=sequence
//ff:what TestExistExcept — two edges against the exception that an adjective's kind does not take what follows: the stem of 없다 stands before 느냐 (the lift's reason in the trace), the stem of 좋다 does not (with that exception's reason), and neither stands before a closing only verbs take (없는다), where the lift is itself defeated with its reason
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestExistExcept(t *testing.T) {
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	tail := except(g, rule.StemTail, whyStemTail, &rule.Bits{Of: "tail", Mask: kindTails[kindAdj]})
	tail.Attacks(claim)
	if n := existExcept(g, tail); n != 2 {
		t.Fatal("edges", n)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "없", Tail: rule.TailExist}, "", true); v <= 0 || !slices.Contains(why, whyExistStem) {
		t.Fatal("없 before 느냐", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "좋", Tail: rule.TailExist}, "", true); v > 0 || !slices.Contains(why, whyStemTail) {
		t.Fatal("좋 before 느냐", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "없", Tail: rule.TailVerb}, "", true); v > 0 || !slices.Contains(why, whyStemTail) || !slices.Contains(why, whyNotExist) {
		t.Fatal("없 before a closing only verbs take", v, why)
	}
	if v, _ := verdictOf(g, rule.Piece{Form: "좋", Tail: rule.TailPred}, "", true); v <= 0 {
		t.Fatal("좋 before another predicate tail", v)
	}
}
