//ff:func feature=parse type=rule control=sequence
//ff:what TestBoundStem — applies to a word that begins with a dependent noun
package parse

import "testing"

func TestBoundStem(t *testing.T) {
	yes, _ := boundStem(ctxFor(readWord("것을", false, &sentence{}), tailFact{}), nil)
	no, _ := boundStem(ctxFor(readWord("밥을", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
