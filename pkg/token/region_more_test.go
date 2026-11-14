//ff:func feature=token type=helper control=sequence
//ff:what TestRegionMore — the ways of the region come in the order of lambda, greatest first, the one that is only a joining of words (the best, at -2.5) left out and counted, and so is the last (-5): the ways of cutting are [0-1][1-2][2-3][3-5] (-3.5) and then [0-2][2-3][3-4][4-5] (-4), and then there is none (and it stays so)
package token

import "testing"

func TestRegionMore(t *testing.T) {
	rg := skewRegion()
	a, ok := rg.more()
	if !ok || segmentLine(a) != "가[0-1] 나[1-2] 다[2-3] 라마[3-5]" || a.Lambda != -3.5 || rg.skipped != 1 {
		t.Fatal(a, ok, rg.skipped)
	}
	rg.out = append(rg.out, a)
	b, ok := rg.more()
	if !ok || segmentLine(b) != "가나[0-2] 다[2-3] 라[3-4] 마[4-5]" || b.Lambda != -4 || b.Rank != 2 {
		t.Fatal(b, ok)
	}
	if _, ok := rg.more(); ok || rg.skipped != 2 {
		t.Fatal(ok, rg.skipped)
	}
	if _, ok := rg.more(); ok {
		t.Fatal("none is left")
	}
}
