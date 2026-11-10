//ff:func feature=mini type=rule control=sequence
//ff:what TestVocativeCase — 철수야 is a call; 철수가 is not
package mini

import "testing"

func TestVocativeCase(t *testing.T) {
	yes, _ := vocativeCase(ctxFor(readWord("철수야", false, &sentence{}), tailFact{}), nil)
	no, _ := vocativeCase(ctxFor(readWord("철수가", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
