//ff:func feature=rule type=helper control=sequence
//ff:what TestShare — the rival's part of the two counts on [-1, 1]; no value when either count is unknown
package rule

import "testing"

func TestShare(t *testing.T) {
	if Share(1, 3) != 0.5 || Share(5, 5) != 0 || Share(0, 3) != -1 || Share(3, 0) != -1 {
		t.Fatal("Share")
	}
}
