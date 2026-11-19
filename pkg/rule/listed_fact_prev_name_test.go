//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFactPrevName — ListPrevName is the whole name of the stem of the nearest word before the candidate (열쇠:동사 of 관하+여, not its kind alone and not the farther words); "" when no word is before it, when that word has no stem name, or when the context has no neighbours; a list of it validates
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactPrevName(t *testing.T) {
	for _, c := range []struct {
		before []Near
		want   string
	}{
		{[]Near{{Stem: "관하", Kind: "열쇠:동사", Last: "맺음:여"}, {Stem: "점", Kind: "열쇠:체언"}}, "열쇠:동사"},
		{[]Near{{Stem: "있", Kind: "보조 용언:있"}}, "보조 용언:있"},
		{[]Near{{Stem: "쪼간"}, {Stem: "먹", Kind: "열쇠:동사"}}, ""},
		{nil, ""},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "못하"}, Before: c.before})
		if got := listedFact(ctx, ListPrevName); got != c.want {
			t.Errorf("%+v: ListPrevName = %q, want %q", c.before, got, c.want)
		}
	}
	if got := listedFact(toulmin.NewContext(), ListPrevName); got != "" {
		t.Errorf("no neighbours: %q", got)
	}
	if err := (&List{Name: "앞 줄기 이름", Of: ListPrevName}).Validate(); err != nil {
		t.Error(err)
	}
}
