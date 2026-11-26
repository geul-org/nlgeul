//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestMarkerGraphCalled — the graph of the suffix 이 read where one is called stands only on a piece told that it ends its unit's letters before a comma or an exclamation mark, and has one edge more than the same entry without the mark; the graph of the suffix read inside a chain stands only with a piece after it
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphCalled(t *testing.T) {
	called := &entry{name: nameSuffix, form: "이", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkSuffix, called: true}
	plain := &entry{name: nameSuffix, form: "이", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkSuffix}
	inner := &entry{name: nameSuffix, form: "이", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkSuffix, inner: true}
	n, m := markerGraph(called, nil, nil, false), markerGraph(plain, nil, nil, false)
	if n != m+1 || markerGraph(inner, nil, nil, false) != n {
		t.Fatal(n, m)
	}
	at := func(next int, call bool) rule.Piece {
		return rule.Piece{Form: "이", HasStem: true, Prev: '숙', Slot: rule.LinkSuffix, Next: next, Called: call}
	}
	for _, c := range []struct {
		e    *entry
		p    rule.Piece
		want bool
	}{
		{called, at(-1, true), true},
		{called, at(-1, false), false},
		{inner, at(-1, true), false},
		{inner, at(-1, false), false},
		{inner, at(rule.LinkCase, false), true},
	} {
		if got := verdict(c.e.g, c.p) > 0; got != c.want {
			t.Errorf("%+v on %+v: %v, want %v", c.e, c.p, got, c.want)
		}
	}
}
