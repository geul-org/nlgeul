//ff:func feature=token type=reader control=sequence
//ff:what TestWholeShown — a word no table lists, followed by letters with no change of script, stands as a word of no markers only when a sign shows the boundary there (shown); otherwise it is held back
package token

import "testing"

func TestWholeShown(t *testing.T) {
	sc := scanOf("구름돌")
	best := [][]*path{{{verdict: 1}}, nil, nil, nil}
	if got := sc.whole(nil, best, 2, false); len(got) != 0 {
		t.Fatal("held back", got)
	}
	got := sc.whole(nil, best, 2, true)
	if len(got) != 1 || got[0].sig != "구름" {
		t.Fatal("shown", got)
	}
}
