//ff:func feature=parse type=rule control=sequence
//ff:what TestChainReadsMore — applies when the chain is longer than the tail under judgment
package parse

import "testing"

func TestChainReadsMore(t *testing.T) {
	w := &word{chain: chainFact{size: 6}}
	yes, _ := chainReadsMore(ctxFor(w, tailFact{size: 2}), nil)
	no, _ := chainReadsMore(ctxFor(w, tailFact{size: 6}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
