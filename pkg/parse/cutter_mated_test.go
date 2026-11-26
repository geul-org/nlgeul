//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCutterMated — a sign written apart has its partner written apart too: an opening < with a > apart after it, a closing > with a < apart before it; not when the sign itself or the other is written onto a word, when the other stands on the wrong side, is of another kind or in a span of code, or when there is none
package parse

import "testing"

func TestCutterMated(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		text string
		at   int
		want bool
	}{
		{"< 가 >", 0, true},
		{"< 가 >", 4, true},
		{"가 < 나 > 다", 2, true},
		{"가 < 나 > 다", 6, true},
		{"<가 >", 0, false},
		{"< 가>", 0, false},
		{"<가 >", 3, false},
		{"가 < 나", 2, false},
		{"가 > 나", 2, false},
		{"> 가 <", 0, false},
		{"> 가 <", 4, false},
		{"< 가 )", 0, false},
		{"( 가 >", 4, false},
		{"< 가 ` > `", 0, false},
		{"< 가 < 나", 0, false},
	}
	for _, c := range cases {
		cut := newCutter(ix, c.text, false)
		if got := cut.mated(c.at, pairRows[cut.r[c.at]]); got != c.want {
			t.Errorf("%q at %d: %v", c.text, c.at, got)
		}
	}
}
