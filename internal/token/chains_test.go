//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestChains — 학교에서는 gives the chains 는 and 에서+는 (and the short 서+는 is refused by order); nothing at the start, nothing past three particles
package token

import "testing"

func TestChains(t *testing.T) {
	sc := scanOf("학교에서는")
	got := sc.chains(5, 0, -1)
	starts := map[int]int{}
	for _, c := range got {
		starts[c.start] = len(c.pieces)
	}
	if starts[4] != 1 || starts[2] != 2 || len(starts) != 2 {
		t.Fatal(starts)
	}
	if sc.chains(0, 0, -1) != nil || sc.chains(5, 3, -1) != nil {
		t.Fatal("limits")
	}
}
