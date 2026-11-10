//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSlots — before a closing: 었 alone and 시+었 in 가시었 (and 었 itself as a syllable with a fused ㅆ); a fused ㅆ in 갔 ends the run, once for each stem it may come from; nothing at the start
package token

import "testing"

func TestSlots(t *testing.T) {
	sc := scanOf("가시었")
	got := sc.slots(3, 9)
	starts := map[int]int{}
	for _, x := range got {
		starts[x.start] = len(x.pieces)
	}
	if starts[2] != 1 || starts[1] != 2 || starts[3] != 1 {
		t.Fatal(got)
	}
	sc = scanOf("갔")
	if got = sc.slots(1, 9); len(got) != 2 || got[0].start != 1 || got[0].pieces[0].text != "(았)" || got[0].base != '가' {
		t.Fatal(got)
	}
	if sc.slots(0, 9) != nil {
		t.Fatal("start")
	}
}
