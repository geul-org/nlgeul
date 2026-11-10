//ff:func feature=token type=reader control=sequence
//ff:what TestWhole — a word of no markers is read at the unit's end from every place, and inside the unit only when it is a closed-set word
package token

import "testing"

func TestWhole(t *testing.T) {
	sc := scanOf("것곰")
	best := [][]*path{{{verdict: 1}}, nil, nil}
	if got := sc.whole(nil, best, 1); len(got) != 1 || got[0].sig != "것[의존명사:것]" {
		t.Fatal(got)
	}
	best[1] = sc.whole(nil, best, 1)
	if got := sc.whole(nil, best, 2); len(got) != 2 {
		t.Fatal(got)
	}
	if got := scanOf("곰것").whole(nil, best[:2], 1); len(got) != 0 {
		t.Fatal(got)
	}
}
