//ff:func feature=token type=reader control=sequence
//ff:what TestScanJoined — a state before a state of the net is added to its preds once: the second time the same edge is given nothing is added, another state before is added
package token

import "testing"

func TestScanJoined(t *testing.T) {
	sc := scanOf("가나")
	state, a, b := &path{}, &path{}, &path{}
	sc.joined(state, a)
	sc.joined(state, a)
	if len(state.preds) != 1 || state.preds[0] != a {
		t.Fatal("the same edge twice", len(state.preds))
	}
	sc.joined(state, b)
	if len(state.preds) != 2 || state.preds[1] != b {
		t.Fatal("another state before", len(state.preds))
	}
}
