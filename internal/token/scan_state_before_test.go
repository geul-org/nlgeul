//ff:func feature=token type=helper control=sequence
//ff:what TestScanStateBefore — a candidate whose word is the first has the scan's root before it, whatever object its root is; any other has the candidate before it
package token

import "testing"

func TestScanStateBefore(t *testing.T) {
	sc := &scan{}
	other := &path{verdict: 1}
	first := grow(other, wordStep(1, "x", 1, -1))
	if sc.stateBefore(first) != sc.registry().root || sc.stateBefore(first) == other {
		t.Fatal("the scan's root")
	}
	second := grow(first, wordStep(2, "y", 1, -1))
	if sc.stateBefore(second) != first {
		t.Fatal("the candidate before")
	}
}
