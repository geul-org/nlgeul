//ff:func feature=rule type=rule control=sequence
//ff:what TestChainSeenMore — the particle's share of the counts when the chain is as long as the tail; no value for a shorter chain or an unknown count
package rule

import "testing"

func TestChainSeenMore(t *testing.T) {
	w := &Word{Chain: Chain{Size: 2, Last: "은", Count: 3}}
	v, ev := ChainSeenMore(ctxFor(w, Tail{Size: 2, Count: 1}), nil)
	short, _ := ChainSeenMore(ctxFor(w, Tail{Size: 4, Count: 1}), nil)
	unknown, _ := ChainSeenMore(ctxFor(w, Tail{Size: 2}), nil)
	if v != 0.5 || ev != "은" || short != -1 || unknown != -1 {
		t.Fatal(v, ev, short, unknown)
	}
}
