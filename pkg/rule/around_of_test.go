//ff:func feature=rule type=helper control=sequence
//ff:what TestAroundOf — the candidate with the words around it set in the context comes back whole; a context without one, or with something else under KeyAround, gives an empty one
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAroundOf(t *testing.T) {
	a := Around{Self: Near{Stem: "학생", Last: "보격:이"}, Before: []Near{{Stem: "밥"}}, After: []Near{{Stem: "되", Pred: true}, {Stem: "다"}}}
	ctx := toulmin.NewContext()
	ctx.Set(KeyAround, a)
	got := AroundOf(ctx)
	if got.Self != a.Self || len(got.Before) != 1 || got.Before[0].Stem != "밥" || len(got.After) != 2 || !got.After[0].Pred || got.After[1].Stem != "다" {
		t.Fatal(got)
	}
	if e := AroundOf(toulmin.NewContext()); e.Self != (Near{}) || e.Before != nil || e.After != nil {
		t.Fatal("empty context", e)
	}
	wrong := toulmin.NewContext()
	wrong.Set(KeyAround, "not an Around")
	if e := AroundOf(wrong); e.Self != (Near{}) || e.Before != nil {
		t.Fatal("another type under the key", e)
	}
}
