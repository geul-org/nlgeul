//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestMarkIsNot — with Not it applies when the sign has none of the facts asked for (one dot alone, a name not of digits: not the shape of a version), and not when it has one (a version, a three-run number, a decimal)
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkIsNot(t *testing.T) {
	cases := []struct {
		m    Mark
		mask int
		want float64
	}{
		{Mark{Run: 1}, MarkRepeated, 1},
		{Mark{Run: 3}, MarkRepeated, -1},
		{Mark{Numeric: true, Parts: []string{"v1", "2"}}, MarkNumeric, -1},
		{Mark{Numeric: true, Parts: []string{"1", "2", "3"}}, MarkNumeric, -1},
		{Mark{Numeric: true, Parts: []string{"1", "2"}}, MarkNumeric, -1},
		{Mark{Parts: []string{"a", "b", "c"}}, MarkNumeric, 1},
		{Mark{}, MarkAfterWord, 1},
		{Mark{Left: SideHangul}, MarkAfterWord, -1},
	}
	for _, c := range cases {
		got, ev := MarkIs(markCtx(c.m), toulmin.Specs{&Bits{Of: "mark", Mask: c.mask}, &Not{}, &Claim{Name: "a", Concludes: "b"}})
		if got != c.want || ev != nil {
			t.Errorf("MarkIs(%+v, %b, not) = %v, %v", c.m, c.mask, got, ev)
		}
	}
}
