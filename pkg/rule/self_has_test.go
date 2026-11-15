//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestSelfHas — applies (+1, with the name as evidence) when one of the candidate's markers is in the list, whichever place it has among them; not (-1, no evidence) when none is, when the candidate has no marker, or when the context holds no candidate; Not turns both round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSelfHas(t *testing.T) {
	l := &List{Name: "adverbial", Of: ListSelfMarks, Items: []string{"부사격:에서", "부사격:서"}}
	ctxOf := func(marks string) toulmin.Context {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "학교", Marks: marks}})
		return ctx
	}
	cases := []struct {
		name string
		ctx  toulmin.Context
		not  bool
		want float64
		fact any
	}{
		{"alone", ctxOf("부사격:에서"), false, 1, "부사격:에서"},
		{"second", ctxOf("보조사:도+부사격:서"), false, 1, "부사격:서"},
		{"first of two", ctxOf("부사격:에서+보조사:는"), false, 1, "부사격:에서"},
		{"other marker", ctxOf("주격:에서"), false, -1, nil},
		{"only a part of a name", ctxOf("부사격:에"), false, -1, nil},
		{"no marker", ctxOf(""), false, -1, nil},
		{"no candidate", toulmin.NewContext(), false, -1, nil},
		{"not, held", ctxOf("부사격:서"), true, -1, "부사격:서"},
		{"not, missing", ctxOf("주격:에서"), true, 1, nil},
	}
	for _, c := range cases {
		specs := toulmin.Specs{l}
		if c.not {
			specs = toulmin.Specs{l, &Not{}}
		}
		got, fact := SelfHas(c.ctx, specs)
		if got != c.want || fact != c.fact {
			t.Errorf("%s: SelfHas = %v, %v; want %v, %v", c.name, got, fact, c.want, c.fact)
		}
	}
}
