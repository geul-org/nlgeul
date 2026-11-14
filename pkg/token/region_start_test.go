//ff:func feature=token type=helper control=sequence
//ff:what TestRegionStart — the words that begin the region go in the frontier, one walk each of f the most it can reach, the better first; words that begin later are not put in
package token

import "testing"

func TestRegionStart(t *testing.T) {
	rg := skewRegion()
	rg.front = frontier{}
	rg.start()
	if len(rg.front.items) != 2 || rg.front.items[0].node != 1 || rg.front.items[0].f != -2.5 || rg.front.items[1].node != 0 || rg.front.items[1].f != -3.5 || rg.front.items[0].prev != nil {
		t.Fatal(rg.front.items)
	}
}
