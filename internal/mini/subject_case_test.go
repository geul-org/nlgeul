//ff:func feature=mini type=rule control=sequence
//ff:what TestSubjectCase — 곰이 and 정부에서 are read as subjects; 문장을 and 매우 are not
package mini

import "testing"

func TestSubjectCase(t *testing.T) {
	yes, ev := subjectCase(ctxFor(readWord("곰이", false, &sentence{}), tailFact{}), nil)
	also, _ := subjectCase(ctxFor(readWord("정부에서", false, &sentence{}), tailFact{}), nil)
	no, _ := subjectCase(ctxFor(readWord("문장을", false, &sentence{}), tailFact{}), nil)
	none, _ := subjectCase(ctxFor(readWord("매우", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || ev != "이" || also != 1 || no != -1 || none != -1 {
		t.Fatal(yes, ev, also, no, none)
	}
}
