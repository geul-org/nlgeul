//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestMarkIs — applies when the sign has one of the facts asked for (two dots or more in a row, a letter on its left, a name of digits); not when it has none of them, or when the context has no sign
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkIs(t *testing.T) {
	cases := []struct {
		m    Mark
		mask int
		want float64
	}{
		{Mark{Run: 3}, MarkRepeated, 1},
		{Mark{Run: 1}, MarkRepeated, -1},
		{Mark{Left: SideHangul}, MarkAfterWord, 1},
		{Mark{}, MarkAfterWord, -1},
		{Mark{Numeric: true, Parts: []string{"v1", "2", "3"}}, MarkNumeric, 1},
		{Mark{Ext: "b", Last: "com", Parts: []string{"a", "b", "com"}}, MarkNumeric, -1},
		{Mark{Left: SideAlnum, Run: 1}, MarkRepeated | MarkAfterWord, 1},
		{Mark{Left: SideDigit, Numeric: true, Run: 1}, MarkRepeated, -1},
	}
	for _, c := range cases {
		got, ev := MarkIs(markCtx(c.m), toulmin.Specs{&Bits{Of: "mark", Mask: c.mask}})
		if got != c.want || ev != nil {
			t.Errorf("MarkIs(%+v, %b) = %v, %v", c.m, c.mask, got, ev)
		}
	}
	if empty, _ := MarkIs(toulmin.NewContext(), toulmin.Specs{&Bits{Of: "mark", Mask: MarkRepeated | MarkAfterWord | MarkNumeric}}); empty != -1 {
		t.Error("empty context", empty)
	}
}
