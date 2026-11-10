//ff:func feature=mini type=rule control=sequence
//ff:what TestFinalWord — applies to the last word only
package mini

import "testing"

func TestFinalWord(t *testing.T) {
	yes, _ := finalWord(ctxFor(&word{last: true}, tailFact{}), nil)
	no, _ := finalWord(ctxFor(&word{}, tailFact{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
