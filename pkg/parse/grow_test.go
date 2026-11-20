//ff:func feature=parse type=reader control=sequence
//ff:what TestGrow — the least verdict and the summed preference; the words printed in order; a word after an 어/아 form sums the preference as any other
package parse

import "testing"

func TestGrow(t *testing.T) {
	p := grow(&path{verdict: 1}, step{sig: "곰+이", verdict: 0.5, pref: -2})
	p = grow(p, step{sig: "간+다", verdict: 0.8, pref: -3})
	if p.sig != "곰+이 간+다" || p.verdict != 0.5 || p.pref != -5 || p.prev.sig != "곰+이" {
		t.Fatal(p)
	}
	aux := step{sig: "보+다", verdict: 1, pref: -4}
	if after := grow(grow(&path{verdict: 1}, step{sig: "먹+어", verdict: 1, pref: -2, conn: true}), aux); after.pref != -6 {
		t.Fatal("after an 어/아 form", after.pref)
	}
}
