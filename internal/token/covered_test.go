//ff:func feature=token type=helper control=sequence
//ff:what TestCovered — the places from an opening sign to its closing one are covered, the others not
package token

import "testing"

func TestCovered(t *testing.T) {
	got := covered(map[int]int{1: 3}, 5)
	if got[0] || !got[1] || !got[2] || !got[3] || got[4] {
		t.Fatal(got)
	}
}
