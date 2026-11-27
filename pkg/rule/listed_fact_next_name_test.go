//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFactNextName — ListNextName is the whole name of the stem of the nearest word after the candidate (원:수, not its kind alone and not the farther words); "" when no word is after it or that word has no stem name; a list of it validates, one with no name is refused
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
		{[]Near{{Stem: "①", Kind: "원:수"}, {Stem: "점", Kind: "열쇠:체언"}}, "원:수"},
		{[]Near{{Stem: "쪼간"}, {Stem: "먹", Kind: "열쇠:동사"}}, ""},
		{nil, ""},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "제"}, After: c.after})
		if got := listedFact(ctx, ListNextName); got != c.want {
			t.Errorf("%+v: ListNextName = %q, want %q", c.after, got, c.want)
		}
	}
	if err := (&List{Name: "뒤 줄기 이름", Of: ListNextName}).Validate(); err != nil {
		t.Error(err)
	}
	if err := (&List{Of: ListNextName}).Validate(); err == nil {
		t.Error("a list with no name is refused")
	}
}
