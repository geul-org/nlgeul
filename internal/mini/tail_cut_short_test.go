//ff:func feature=mini type=rule control=sequence
//ff:what TestTailCutShort — applies to a tail shorter than the word's longest
package mini

import "testing"

func TestTailCutShort(t *testing.T) {
	w := &word{longest: 3}
	yes, _ := tailCutShort(ctxFor(w, tailFact{size: 2}), nil)
	no, _ := tailCutShort(ctxFor(w, tailFact{size: 3}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
