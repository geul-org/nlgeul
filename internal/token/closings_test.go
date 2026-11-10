//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosings — 먹었다 gives 다 alone and 었+다; 간다 a closing under the consonant; nothing at the start
package token

import "testing"

func TestClosings(t *testing.T) {
	sc := scanOf("먹었다")
	got := sc.closings(3, false)
	starts := map[int]bool{}
	for _, c := range got {
		starts[c.start] = true
	}
	if !starts[2] || !starts[1] {
		t.Fatal(got)
	}
	sc = scanOf("간다")
	var fused bool
	for _, c := range sc.closings(2, false) {
		fused = fused || c.pieces[0].text == "(ㄴ)다"
	}
	if !fused || sc.closings(0, false) != nil {
		t.Fatal(fused)
	}
}
