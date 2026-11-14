//ff:func feature=token type=helper control=sequence
//ff:what TestRegionPeek — the best way of a region is made when it has not been, is not taken as sent, and is the one deal gives first; a region with no way gives none
package token

import "testing"

func TestRegionPeek(t *testing.T) {
	r, st := windowCut(t, 5)
	_ = r
	rg := newRegionNet(Contest{Text: st.contest.Text}, 0, *st.cells)
	seg, ok := rg.peek()
	if !ok || rg.sent != 0 || len(rg.out) != 1 || seg.Rank != 1 {
		t.Fatal(ok, rg.sent)
	}
	again, _ := rg.peek()
	dealt, _ := rg.deal()
	if again.Rank != 1 || dealt.Rank != 1 || rg.sent != 1 || len(rg.out) != 1 {
		t.Fatal("the same way")
	}
	empty := newRegion(Contest{Text: "가나"}, 0, nil)
	if _, ok := empty.peek(); ok {
		t.Fatal("no way")
	}
}
