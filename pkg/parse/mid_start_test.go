//ff:func feature=parse type=helper control=sequence
//ff:what TestMidStart — the run of one sign that ends before a place begins at its first sign; a single sign begins at itself; a place past the unit's letters gives the place before it
package parse

import "testing"

func TestMidStart(t *testing.T) {
	sc := scanOf("산——뭐")
	if got := sc.midStart(3); got != 1 {
		t.Errorf("run: %d", got)
	}
	if got := sc.midStart(2); got != 1 {
		t.Errorf("first of the run: %d", got)
	}
	if got := sc.midStart(1); got != 0 {
		t.Errorf("head: %d", got)
	}
	if got := sc.midStart(9); got != 8 {
		t.Errorf("past the letters: %d", got)
	}
}
