//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFactPrevKind — ListPrevKind is the kind of the name of the stem of the nearest word before the candidate, cut at the first colon (닫기:1 gives 닫기, a kind with no colon is kept whole); "" when no word is before it, when the name of that stem is empty, or when the context has no neighbours
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactPrevKind(t *testing.T) {
	for _, c := range []struct {
		before []Near
		want   string
	}{
		{[]Near{{Stem: ")", Kind: "닫기:1"}, {Stem: "먼", Kind: "열기:1"}}, "닫기"},
		{[]Near{{Stem: "책", Kind: "열쇠:체언"}}, "열쇠"},
		{[]Near{{Stem: "가", Kind: "닫기"}}, "닫기"},
		{[]Near{{Stem: "새"}}, ""},
		{nil, ""},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "라고"}, Before: c.before})
		if got := listedFact(ctx, ListPrevKind); got != c.want {
			t.Errorf("%+v: ListPrevKind = %q, want %q", c.before, got, c.want)
		}
	}
	if got := listedFact(toulmin.NewContext(), ListPrevKind); got != "" {
		t.Errorf("no neighbours: %q", got)
	}
}
