//ff:func feature=token type=reader control=sequence
//ff:what TestGrow — the least verdict and the summed preference; the words printed in order
package token

import "testing"

func TestGrow(t *testing.T) {
	p := grow(&path{verdict: 1}, eojeol{marked: "곰+이", verdict: 0.5, pref: -2})
	p = grow(p, eojeol{marked: "간+다", verdict: 0.8, pref: -3})
	if p.sig != "곰+이 간+다" || p.verdict != 0.5 || p.pref != -5 || p.prev.sig != "곰+이" {
		t.Fatal(p)
	}
}
