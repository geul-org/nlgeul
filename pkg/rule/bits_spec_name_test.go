//ff:func feature=rule type=model control=sequence
//ff:what TestBitsSpecName — the name carries the set
package rule

import "testing"

func TestBitsSpecName(t *testing.T) {
	if (&Bits{Of: "case"}).SpecName() != "bits:case" {
		t.Fatal("SpecName")
	}
}
