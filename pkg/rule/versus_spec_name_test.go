//ff:func feature=rule type=model control=sequence
//ff:what TestVersusSpecName — the name is versus
package rule

import "testing"

func TestVersusSpecName(t *testing.T) {
	if (&Versus{}).SpecName() != "versus" {
		t.Fatal("SpecName")
	}
}
