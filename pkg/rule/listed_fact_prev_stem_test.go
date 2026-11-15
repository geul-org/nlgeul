//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFactPrevStem — ListPrevStem is the stem of the nearest word before the candidate as printed (관하 of 관하+여, not the farther ones); "" when no word is before it, when that word has no stem, or when the context has no neighbours
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactPrevStem(t *testing.T) {
	for _, c := range []struct {
		before []Near
		want   string
	}{
		{[]Near{{Stem: "관하", Kind: "열쇠:동사", Last: "맺음:여"}, {Stem: "점", Kind: "열쇠:체언"}}, "관하"},
		{[]Near{{Stem: "먹", Kind: "열쇠:동사"}}, "먹"},
		{[]Near{{Last: "쉼표"}, {Stem: "관하"}}, ""},
		{nil, ""},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "보"}, Before: c.before})
		if got := listedFact(ctx, ListPrevStem); got != c.want {
			t.Errorf("%+v: ListPrevStem = %q, want %q", c.before, got, c.want)
		}
	}
	if got := listedFact(toulmin.NewContext(), ListPrevStem); got != "" {
		t.Errorf("no neighbours: %q", got)
	}
}
