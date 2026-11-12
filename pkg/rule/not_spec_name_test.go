//ff:func feature=rule type=model control=sequence
//ff:what TestNotSpecName — the name is not
package rule

import "testing"

func TestNotSpecName(t *testing.T) {
	if (&Not{}).SpecName() != "not" {
		t.Fatal("SpecName")
	}
}
