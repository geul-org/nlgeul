//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveWaysBig — over 100 letters with a word of one letter and a word of two at every place the ways are the Fibonacci number 573147844013817084101, more than a 64-bit integer holds, counted exactly
package token

import "testing"

func TestWeaveWaysBig(t *testing.T) {
	sc, lat := fibNet(100)
	pts := lat.portions()
	if len(pts) != 1 {
		t.Fatal(len(pts))
	}
	w := newWeave(scanOf("x").ix, pts[0])
	got := w.ways()
	if got.String() != "573147844013817084101" || got.IsInt64() || sc == nil {
		t.Fatal(got)
	}
}
