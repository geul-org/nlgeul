//ff:func feature=parse type=reader control=sequence
//ff:what TestTrimI — the copula stem read at the end of a tail's stem; nothing without a copula or when it is fused elsewhere
package parse

import "testing"

func TestTrimI(t *testing.T) {
	if trimI(tailFact{stem: "것이", copula: true}) != "이" || trimI(tailFact{stem: "것이"}) != "" || trimI(tailFact{stem: "학교", copula: true}) != "" {
		t.Fatal("trimI")
	}
}
