//ff:func feature=rule type=helper control=sequence
//ff:what TestInRange — no ranges ask nothing; a value is held to its own range, and past the list to the last; a top of 0 is no bound
package rule

import "testing"

func TestInRange(t *testing.T) {
	r := [][2]int{{4, 4}, {1, 2}}
	if !InRange(nil, 3, 99) || !InRange(r, 0, 4) || InRange(r, 0, 3) || !InRange(r, 5, 2) || InRange(r, 5, 3) || !InRange([][2]int{{1, 0}}, 0, 999) || InRange([][2]int{{1, 0}}, 0, 0) {
		t.Fatal("InRange")
	}
}
