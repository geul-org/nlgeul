//ff:func feature=rule type=model control=sequence
//ff:what TestVersusValidate — both sides must be named
package rule

import "testing"

func TestVersusValidate(t *testing.T) {
	if (&Versus{Mine: SideTail, Rival: SideChain}).Validate() != nil || (&Versus{Mine: SideTail}).Validate() == nil || (&Versus{Rival: SideChain}).Validate() == nil {
		t.Fatal("Validate")
	}
}
