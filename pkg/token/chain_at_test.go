//ff:func feature=token type=reader control=sequence
//ff:what TestChainAt — runs tied by single dots make one dotted name; a run after something else starts another; two runs of digits are tied by a sign of numbers, a letter run is not
package token

import "testing"

func TestChainAt(t *testing.T) {
	s := []rune("a.b.c d.e")
	runs := runsOf(s)
	if got := chainAt(s, runs, 0); len(got) != 3 {
		t.Fatal(got)
	}
	if got := chainAt(s, runs, 3); len(got) != 2 || got[0] != [2]int{6, 7} {
		t.Fatal(got)
	}
	s2 := []rune("10:30 a:1 a-1")
	r2 := runsOf(s2)
	if len(chainAt(s2, r2, 0)) != 2 || len(chainAt(s2, r2, 2)) != 1 || len(chainAt(s2, r2, 4)) != 2 {
		t.Fatal("a sign of numbers ties digits only, a hyphen any two runs")
	}
}
