//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestPrevHas — applies (+1, with the name as evidence) when one of the markers of the word right before the candidate is in the list, whichever place it has among them; not (-1, no evidence) when none is, when that word has no marker, when no word stands before the candidate, or when the context holds no candidate; the candidate's own markers are not read; Not turns both round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPrevHas(t *testing.T) {
	l := &List{Name: "naming", Of: ListPrevLast, Items: []string{"이름:기"}}
	ctxOf := func(marks string) toulmin.Context {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "하", Marks: "이름:기"}, Before: []Near{{Stem: "있", Marks: marks}, {Stem: "돈", Marks: "이름:기"}}})
		return ctx
	}
	alone := toulmin.NewContext()
	alone.Set(KeyAround, Around{Self: Near{Stem: "하", Marks: "이름:기"}})
	cases := []struct {
		name string
		ctx  toulmin.Context
		not  bool
		want float64
		fact any
	}{
		{"first", ctxOf("이름:기+보조사:는"), false, 1, "이름:기"},
		{"alone", ctxOf("이름:기"), false, 1, "이름:기"},
		{"other markers", ctxOf("맺음:고"), false, -1, nil},
		{"no marker", ctxOf(""), false, -1, nil},
		{"no word before", alone, false, -1, nil},
		{"no candidate", toulmin.NewContext(), false, -1, nil},
		{"not, held", ctxOf("이름:기+보조사:도"), true, -1, "이름:기"},
		{"not, missing", ctxOf("맺음:고"), true, 1, nil},
	}
	for _, c := range cases {
		specs := toulmin.Specs{l}
		if c.not {
			specs = toulmin.Specs{l, &Not{}}
		}
		got, fact := PrevHas(c.ctx, specs)
		if got != c.want || fact != c.fact {
			t.Errorf("%s: %v %v, want %v %v", c.name, got, fact, c.want, c.fact)
		}
	}
}
