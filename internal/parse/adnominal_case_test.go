//ff:func feature=parse type=rule control=sequence
//ff:what TestAdnominalCase — 상급자의 and 학교에서의 are adnominal; 학교에서 is not
package parse

import "testing"

func TestAdnominalCase(t *testing.T) {
	yes, _ := adnominalCase(ctxFor(readWord("상급자의", false, &sentence{}), tailFact{}), nil)
	outside, _ := adnominalCase(ctxFor(readWord("학교에서의", false, &sentence{}), tailFact{}), nil)
	no, _ := adnominalCase(ctxFor(readWord("학교에서", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || outside != 1 || no != -1 {
		t.Fatal(yes, outside, no)
	}
}
