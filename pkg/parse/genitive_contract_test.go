//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestGenitiveContract — 내, 네 and 제 read as 나, 너 and 저 with 의 alone are forms of the contracts table; 나의 written out, the form read as a nominal with no particle, the form with another particle and a stem that is no nominal of the table are not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGenitiveContract(t *testing.T) {
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Marks: "관형격:의", Text: "내"}, true},
		{rule.Near{Stem: "너", Kind: "열쇠:체언", Marks: "관형격:의", Text: "네"}, true},
		{rule.Near{Stem: "저", Kind: "열쇠:체언", Marks: "관형격:의", Text: "제"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Marks: "관형격:의", Text: "나의"}, false},
		{rule.Near{Stem: "제", Kind: "열쇠:체언", Text: "제"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Marks: "주격:가", Text: "내가"}, false},
		{rule.Near{Stem: "저", Kind: "관형사:저", Marks: "관형격:의", Text: "제"}, false},
	} {
		if got := genitiveContract(c.near); got != c.want {
			t.Errorf("genitiveContract(%+v) = %v, want %v", c.near, got, c.want)
		}
	}
}
