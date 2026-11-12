//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestMarkFacts — two dots or more in a row set MarkRepeated, anything on the left sets MarkAfterWord, a name of digits sets MarkNumeric; one dot with nothing on the left and a name not of digits has none
package rule

import "testing"

func TestMarkFacts(t *testing.T) {
	cases := []struct {
		m    Mark
		want int
	}{
		{Mark{}, 0},
		{Mark{Run: 1, Right: SideAlnum, Ext: "py", Last: "py", Parts: []string{"a", "py"}}, 0},
		{Mark{Run: 2}, MarkRepeated},
		{Mark{Run: 3}, MarkRepeated},
		{Mark{Left: SideHangul, Run: 1}, MarkAfterWord},
		{Mark{Left: SideAlnum | SideDigit}, MarkAfterWord},
		{Mark{Numeric: true, Run: 1}, MarkNumeric},
		{Mark{Left: SideDigit, Right: SideDigit, Numeric: true, Run: 1}, MarkAfterWord | MarkNumeric},
		{Mark{Left: SideHangul, Run: 3}, MarkRepeated | MarkAfterWord},
	}
	for _, c := range cases {
		if got := markFacts(c.m); got != c.want {
			t.Errorf("markFacts(%+v) = %b, want %b", c.m, got, c.want)
		}
	}
}
