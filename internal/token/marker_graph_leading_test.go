//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMarkerGraphLeading — the auxiliary particles 이야 and 야 (leading) get one more exception, which turns them down when they stand right on a stem and a sign that ends a sentence closes the unit; another particle has no such exception and stands there
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphLeading(t *testing.T) {
	cases := []struct {
		e      entry
		edges  int
		closed float64
	}{
		{entry{name: "보조사:이야", form: "이야", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkTopic}, 4, -1},
		{entry{name: "보조사:야", form: "야", class: classParticle, after: rule.AfterVowel, slot: rule.LinkTopic}, 4, -1},
		{entry{name: "호격:야", form: "야", class: classParticle, after: rule.AfterVowel, slot: rule.LinkVoc}, 3, 1},
		{entry{name: "보조사:도", form: "도", class: classParticle, slot: rule.LinkTopic}, 2, 1},
	}
	for _, c := range cases {
		e := c.e
		if got := markerGraph(&e, nil, false); got != c.edges {
			t.Errorf("%s: %d edges, want %d", e.name, got, c.edges)
		}
		prev := '교'
		if e.after == rule.AfterBatchim {
			prev = '당'
		}
		p := rule.Piece{Form: e.form, HasStem: true, Prev: prev, Slot: e.slot, Next: -1}
		if v := verdict(e.g, p); v != 1 {
			t.Errorf("%s with no sign after it: %v", e.name, v)
		}
		p.Closed = true
		if v := verdict(e.g, p); v != c.closed {
			t.Errorf("%s before a sign that ends a sentence: %v, want %v", e.name, v, c.closed)
		}
	}
}
