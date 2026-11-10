//ff:func feature=token type=helper control=sequence
//ff:what TestWithGate — no graph leaves the tails as they are; a graph is set on each
package token

import "testing"

func TestWithGate(t *testing.T) {
	g := &entry{name: "g"}
	if got := withGate([]tail{{start: 1}}, nil); got[0].gate != nil {
		t.Fatal(got)
	}
	if got := withGate([]tail{{start: 1}, {start: 2}}, g); got[0].gate != g || got[1].gate != g {
		t.Fatal(got)
	}
}
