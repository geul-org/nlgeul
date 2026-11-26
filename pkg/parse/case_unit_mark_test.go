//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCaseUnitMark — the letters between the two numbers of a case number's shape are a unit there when they are a word of the unit table that is no case mark (초, 시, 승, 시간), or a unit that is a case mark too with a piece read after a number right after the second number (36도5분); a case mark that is a unit with nothing after the second number (36도5, 19차3), a greater number and letters no table reads after a number (다) are not
package parse

import "testing"

func TestCaseUnitMark(t *testing.T) {
	cases := []struct {
		text string
		mark string
		end  int
		want bool
	}{
		{"11초92", "초", 5, true},
		{"10시30", "시", 5, true},
		{"81승53", "승", 5, true},
		{"1시간20", "시간", 5, true},
		{"36도5분", "도", 4, true},
		{"36도5", "도", 4, false},
		{"19차3", "차", 4, false},
		{"17만2000", "만", 7, false},
		{"17만2000원", "만", 7, true},
		{"2004다25123", "다", 10, false},
	}
	for _, c := range cases {
		sc := &scan{ix: scanOf("").ix, s: []rune(c.text)}
		if got := sc.caseUnitMark(c.mark, c.end); got != c.want {
			t.Errorf("%s (%s): %v, want %v", c.text, c.mark, got, c.want)
		}
	}
}
