//ff:func feature=token type=helper control=sequence
//ff:what TestRegionTake — a walk to the end of the region is a segment when its spans are one of the ways of cutting; one made by joining words of two ways gives nothing and is counted
package token

import "testing"

func TestRegionTake(t *testing.T) {
	rg := skewRegion()
	seg, ok := rg.take(chainWalk(1, 3, 4, 6))
	if !ok || seg.Lambda != -4 || rg.skipped != 0 {
		t.Fatal(seg, ok, rg.skipped)
	}
	if seg, ok := rg.take(chainWalk(1, 3, 5)); ok || len(seg.Nodes) != 0 || rg.skipped != 1 {
		t.Fatal(seg, ok, rg.skipped)
	}
}
