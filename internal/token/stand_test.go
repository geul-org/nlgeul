//ff:func feature=token type=reader control=sequence
//ff:what TestStand — the letters from b up to j may stand as a word of no markers: the first for a place makes its set, the next only adds to it
package token

import "testing"

func TestStand(t *testing.T) {
	sc := scanOf("뚜루뚜겸")
	sc.stands = make([]map[int]bool, 5)
	sc.stand(4, 3)
	sc.stand(4, 0)
	if len(sc.stands[4]) != 2 || !sc.stands[4][3] || !sc.stands[4][0] || sc.stands[3] != nil {
		t.Fatal(sc.stands)
	}
}
