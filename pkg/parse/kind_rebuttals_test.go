//ff:func feature=parse type=rule control=sequence
//ff:what TestKindRebuttals — six edges, the rebuttals "rarely seen" and "one syllable long" against a claim and what lifts each: on a graph with them a form of more syllables seen enough times stands fully; one syllable long it stands lower, and fully as a sole word or as an auxiliary's stem written right after the 어/아 form of its main verb (Conn), which does not lift "rarely seen"; seen once it stands lower, also as a sole word; with both rebuttals a form stands below either one alone and above 0
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestKindRebuttals(t *testing.T) {
	enough := 10000
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	if n := kindRebuttals(g, claim, enough); n != 6 {
		t.Fatal("edges", n)
	}
	judge := func(form string, count int, sole bool) float64 {
		v, _ := verdictOf(g, rule.Piece{Form: form, Count: count, Sole: sole}, "", false)
		return v
	}
	conn := func(form string, count int) float64 {
		v, _ := verdictOf(g, rule.Piece{Form: form, Count: count, Conn: true}, "", false)
		return v
	}
	if v := judge("가르", enough, false); v != 1 {
		t.Fatal("a form of more syllables seen enough", v)
	}
	short := judge("솥", enough, false)
	if short >= 1 || judge("솥", enough, true) != 1 {
		t.Fatal("a one-syllable form seen enough", short)
	}
	few := judge("가르", 1, false)
	if few >= 1 {
		t.Fatal("a form of more syllables seen once", few)
	}
	if v := conn("솥", enough); v != 1 {
		t.Fatal("a one-syllable auxiliary's stem seen enough right after 어/아", v)
	}
	if v := conn("솥", 1); v != few {
		t.Fatal("seen once right after 어/아 it is still rarely seen", v, few)
	}
	if v := judge("가르", 1, true); v != few {
		t.Fatal("seen once as a sole word it is still rarely seen", v, few)
	}
	both := judge("솥", 1, false)
	if both <= 0 || both >= short || both >= few {
		t.Fatal("a one-syllable form seen once", both, short, few)
	}
}
