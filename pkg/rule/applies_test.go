//ff:func feature=rule type=helper control=sequence
//ff:what TestApplies — yes is +1, no is -1
package rule

import "testing"

func TestApplies(t *testing.T) {
	if Applies(true) != 1 || Applies(false) != -1 {
		t.Fatal("Applies")
	}
}
