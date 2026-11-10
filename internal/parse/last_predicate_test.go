//ff:func feature=parse type=rule control=sequence
//ff:what TestLastPredicate — any tail on the last word makes a predicate; not on an earlier word, not without a tail
package parse

import "testing"

func TestLastPredicate(t *testing.T) {
	tail := tailFact{closing: "는", kind: kindAdn, size: 2}
	yes, _ := lastPredicate(ctxFor(&word{last: true}, tail), nil)
	mid, _ := lastPredicate(ctxFor(&word{}, tail), nil)
	none, _ := lastPredicate(ctxFor(&word{last: true}, tailFact{}), nil)
	if yes != 1 || mid != -1 || none != -1 {
		t.Fatal(yes, mid, none)
	}
}
