//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCueGroups — a Hangul numeral that begins at a place is read in groups at 만, 억 and 조 and each group stands (삼억오천만원, 삼억오천원, 삼억오); one group alone, no numeral, two digits side by side and a place right after a numeral letter make none
package token

import "testing"

func TestCueGroups(t *testing.T) {
	groups := func(text string, at int) []map[int]bool {
		sc := scanOf(text)
		sc.stands = make([]map[int]bool, len(sc.s)+1)
		sc.cueGroups(at)
		return sc.stands
	}
	if got := groups("삼억오천만원을", 0); !got[2][0] || !got[5][2] || len(got[7]) != 0 {
		t.Fatal(got)
	}
	if got := groups("삼억오천원", 0); !got[2][0] || !got[4][2] {
		t.Fatal("a last group with no greater number", got)
	}
	if got := groups("삼억오", 0); !got[2][0] || !got[3][2] {
		t.Fatal("a last group at the end", got)
	}
	for _, c := range []struct {
		text string
		at   int
	}{{"오천만원을", 0}, {"원을", 0}, {"이이억삼", 0}, {"삼억오천만원", 1}} {
		if !standsEmpty(groups(c.text, c.at)) {
			t.Fatal(c.text, groups(c.text, c.at))
		}
	}
}
