//ff:func feature=parse type=rule control=sequence
//ff:what TestObjectCase — 문장을 and 일만을 are objects; 곰이 is not
package parse

import "testing"

func TestObjectCase(t *testing.T) {
	yes, ev := objectCase(ctxFor(readWord("문장을", false, &sentence{}), tailFact{}), nil)
	chain, _ := objectCase(ctxFor(readWord("일만을", false, &sentence{}), tailFact{}), nil)
	no, _ := objectCase(ctxFor(readWord("곰이", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || ev != "을" || chain != 1 || no != -1 {
		t.Fatal(yes, ev, chain, no)
	}
}
