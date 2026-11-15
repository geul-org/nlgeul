//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFactNextName — ListNextName is the whole name of the stem of the nearest word after the candidate (열쇠:동사, 닫기:1 kept whole); "" when no word is after it, when that stem has no name (a new word), or when the context has no neighbours
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactNextName(t *testing.T) {
	for _, c := range []struct {
		after []Near
		want  string
	}{
		{[]Near{{Stem: "하", Kind: "열쇠:동사"}, {Stem: "책", Kind: "열쇠:체언"}}, "열쇠:동사"},
		{[]Near{{Stem: ")", Kind: "닫기:1"}}, "닫기:1"},
		{[]Near{{Stem: "곰곰"}}, ""},
		{nil, ""},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "안", Kind: "열쇠:부사"}, After: c.after})
		if got := listedFact(ctx, ListNextName); got != c.want {
			t.Errorf("%+v: ListNextName = %q, want %q", c.after, got, c.want)
		}
	}
	if got := listedFact(toulmin.NewContext(), ListNextName); got != "" {
		t.Errorf("no neighbours: %q", got)
	}
}
