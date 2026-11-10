//ff:func feature=parse type=rule control=sequence
//ff:what TestAdnominalClosing — a modifying closing makes an adnominal; a final one does not
package parse

import "testing"

func TestAdnominalClosing(t *testing.T) {
	yes, _ := adnominalClosing(ctxFor(&word{}, tailFact{closing: "는", kind: kindAdn, size: 2}), nil)
	no, _ := adnominalClosing(ctxFor(&word{}, tailFact{closing: "다", kind: kindFinal, size: 2}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
