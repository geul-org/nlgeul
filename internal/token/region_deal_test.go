//ff:func feature=token type=helper control=sequence
//ff:what TestRegionDeal — the ways are made one at a time as they are dealt, one after the other; none is made ahead; at the end false, and the count of those sent does not move
package token

import "testing"

func TestRegionDeal(t *testing.T) {
	rg := skewRegion()
	a, ok := rg.deal()
	if !ok || a.Rank != 1 || len(rg.out) != 1 || rg.sent != 1 {
		t.Fatal(a, ok, len(rg.out))
	}
	b, ok := rg.deal()
	if !ok || b.Rank != 2 || len(rg.out) != 2 || rg.sent != 2 {
		t.Fatal(b, ok)
	}
	if _, ok := rg.deal(); ok || rg.sent != 2 || len(rg.out) != 2 {
		t.Fatal("none is left", ok, rg.sent)
	}
	rg.sent = 0
	if again, ok := rg.deal(); !ok || again.Rank != 1 || len(rg.out) != 2 {
		t.Fatal("a way made is given as it was made", again, ok)
	}
}
