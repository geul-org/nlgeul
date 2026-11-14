//ff:func feature=token type=helper control=sequence
//ff:what TestRegionExtend — a walk taken one word further for each word after its last: f is the lambdas up to the last plus the most that can be reached from the new last
package token

import "testing"

func TestRegionExtend(t *testing.T) {
	rg := skewRegion()
	rg.front = frontier{}
	rg.extend(&walk{node: 3, got: -2})
	if len(rg.front.items) != 2 || rg.front.items[0].node != 5 || rg.front.items[0].f != -3.5 || rg.front.items[0].got != -3 || rg.front.items[1].node != 4 || rg.front.items[1].f != -5 || rg.front.items[1].prev == nil {
		t.Fatal(rg.front.items)
	}
	rg.front = frontier{}
	rg.extend(&walk{node: 6})
	if len(rg.front.items) != 0 {
		t.Fatal("the last word has none after it")
	}
}
