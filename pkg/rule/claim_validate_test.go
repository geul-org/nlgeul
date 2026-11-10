//ff:func feature=rule type=model control=sequence
//ff:what TestClaimValidate — a claim needs both a name and a conclusion
package rule

import "testing"

func TestClaimValidate(t *testing.T) {
	if (&Claim{Name: "a", Concludes: "b"}).Validate() != nil || (&Claim{Name: "a"}).Validate() == nil || (&Claim{Concludes: "b"}).Validate() == nil {
		t.Fatal("Validate")
	}
}
