//ff:func feature=parse type=rule control=sequence
//ff:what TestChainSeenMore — the particle's share of the counts when the chain is as long as the tail; no value for a shorter chain or an unknown count
package parse

import "testing"

func TestChainSeenMore(t *testing.T) {
	w := &word{chain: chainFact{size: 2, last: "은", count: 3}}
	v, ev := chainSeenMore(ctxFor(w, tailFact{size: 2, count: 1}), nil)
	short, _ := chainSeenMore(ctxFor(w, tailFact{size: 4, count: 1}), nil)
	unknown, _ := chainSeenMore(ctxFor(w, tailFact{size: 2}), nil)
	if v != 0.5 || ev != "은" || short != -1 || unknown != -1 {
		t.Fatal(v, ev, short, unknown)
	}
}
