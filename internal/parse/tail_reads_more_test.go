//ff:func feature=parse type=rule control=sequence
//ff:what TestTailReadsMore — applies when the longest tail is longer than the chain
package parse

import "testing"

func TestTailReadsMore(t *testing.T) {
	yes, _ := tailReadsMore(ctxFor(&word{chain: chainFact{size: 2}, longest: 4}, tailFact{}), nil)
	no, _ := tailReadsMore(ctxFor(&word{chain: chainFact{size: 2}, longest: 2}, tailFact{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
