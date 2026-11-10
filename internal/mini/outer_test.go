//ff:func feature=mini type=reader control=sequence
//ff:what TestOuter — the right particle's case when it has one, else the left one's
package mini

import "testing"

func TestOuter(t *testing.T) {
	if outer(caseAdn, caseAdv) != caseAdn || outer(0, caseAdv) != caseAdv || outer(0, 0) != 0 {
		t.Fatal("outer")
	}
}
