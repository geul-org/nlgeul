//ff:func feature=token type=reader control=sequence
//ff:what TestCueList — the words of a list are cues at a place when the rule of their article is on, and none when it is off
package token

import "testing"

func TestCueList(t *testing.T) {
	on := spacingScan("뚜루뚜겸총무", "45")
	on.stands, on.shows = make([]map[int]bool, 7), map[int]map[int]bool{}
	on.cueList(3, "45", []string{"겸", "및"})
	if !on.stands[4][3] || !on.shows[3][4] {
		t.Fatal(on.stands, on.shows)
	}
	off := spacingScan("뚜루뚜겸총무", "41")
	off.stands, off.shows = make([]map[int]bool, 7), map[int]map[int]bool{}
	off.cueList(3, "45", []string{"겸"})
	if len(off.shows) != 0 || off.stands[4] != nil {
		t.Fatal("the rule is off", off.stands, off.shows)
	}
}
