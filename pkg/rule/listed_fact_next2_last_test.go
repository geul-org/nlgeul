//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestListedFactNext2Last — ListNext2Last is the name of the last piece of the second word after the candidate (꾸밈:ㄴ of 크+ㄴ after 가장, not that of the nearest word); "" when there is no such word, when that word is its stem alone, or when the context has no neighbours; a list of it validates
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactNext2Last(t *testing.T) {
	for _, c := range []struct {
		after []Near
		want  string
	}{
		{[]Near{{Stem: "가장", Kind: "열쇠:부사"}, {Stem: "크", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ"}}, "꾸밈:ㄴ"},
		{[]Near{{Stem: "가장", Kind: "열쇠:부사", Last: "꾸밈:ㄴ"}, {Stem: "집", Kind: "열쇠:체언"}}, ""},
		{[]Near{{Stem: "가장", Kind: "열쇠:부사"}}, ""},
		{nil, ""},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Stem: "있", Last: "꾸밈:는"}, After: c.after})
		if got := listedFact(ctx, ListNext2Last); got != c.want {
			t.Errorf("%+v: ListNext2Last = %q, want %q", c.after, got, c.want)
		}
	}
	if got := listedFact(toulmin.NewContext(), ListNext2Last); got != "" {
		t.Errorf("no neighbours: %q", got)
	}
	if err := (&List{Name: "그다음 어절의 끝 조각", Of: ListNext2Last}).Validate(); err != nil {
		t.Error(err)
	}
}
