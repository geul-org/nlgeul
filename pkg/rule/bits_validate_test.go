//ff:func feature=rule type=model control=sequence
//ff:what TestBitsValidate — a criterion without a flag is refused
package rule

import "testing"

func TestBitsValidate(t *testing.T) {
	if (&Bits{Mask: 1}).Validate() != nil || (&Bits{}).Validate() == nil {
		t.Fatal("Validate")
	}
}
