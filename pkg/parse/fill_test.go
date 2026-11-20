//ff:func feature=parse type=reader control=sequence
//ff:what TestFill — the places strictly inside are marked, the edges are not
package parse

import "testing"

func TestFill(t *testing.T) {
	sc := &scan{inside: make([]bool, 6)}
	sc.fill(1, 4)
	if sc.inside[1] || !sc.inside[2] || !sc.inside[3] || sc.inside[4] {
		t.Fatal(sc.inside)
	}
}
