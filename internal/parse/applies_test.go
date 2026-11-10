//ff:func feature=parse type=helper control=sequence
//ff:what TestApplies — yes is +1, no is -1
package parse

import "testing"

func TestApplies(t *testing.T) {
	if applies(true) != 1 || applies(false) != -1 {
		t.Fatal("applies")
	}
}
