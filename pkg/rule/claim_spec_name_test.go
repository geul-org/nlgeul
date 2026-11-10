//ff:func feature=rule type=model control=sequence
//ff:what TestClaimSpecName — the name is claim
package rule

import "testing"

func TestClaimSpecName(t *testing.T) {
	if (&Claim{}).SpecName() != "claim" {
		t.Fatal("SpecName")
	}
}
