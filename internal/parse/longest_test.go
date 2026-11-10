//ff:func feature=parse type=reader control=sequence
//ff:what TestLongest — the greatest length among the tails; 0 without tails
package parse

import "testing"

func TestLongest(t *testing.T) {
	if longest([]tailFact{{size: 2}, {size: 5}, {size: 3}}) != 5 || longest(nil) != 0 {
		t.Fatal("longest")
	}
}
