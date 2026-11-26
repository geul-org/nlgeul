//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSignGraphOf — the graph of a candidate is that of its name, but for the hyphen and the box-drawing line the graph that reads their place: the dash ― stands in a line of signs alone (no exception), ─ and - do not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSignGraphOf(t *testing.T) {
	m := rule.Mark{Run: 3, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkBare}
	for sign, want := range map[rune]struct {
		edges int
		stand bool
	}{'―': {0, true}, '—': {0, true}, '─': {2, false}, '-': {7, false}} {
		g, edges := signGraphOf(sign, "줄표")
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, _ := evaluate(g, ctx, false)
		if edges != want.edges || (v > 0) != want.stand {
			t.Errorf("%c: %d edges, verdict %.2f, want %d edges, stands %v", sign, edges, v, want.edges, want.stand)
		}
	}
	if _, edges := signGraphOf(',', "쉼표"); edges != 1 {
		t.Fatal("a sign that is no plain dash has the graph of its name", edges)
	}
}
