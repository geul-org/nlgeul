//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunLines — the spans along the best way to a state, first span first: the line of a state ends at the state's span and begins at the last cut
package token

import "testing"

func TestNetRunLines(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	front, ok := r.front()
	if !ok {
		t.Fatal("no front")
	}
	lines := r.lines(front)
	if len(lines) != len(front) {
		t.Fatal(len(lines))
	}
	for i, ln := range lines {
		last := ln[len(ln)-1]
		if ln[0].Start != 0 || last != (Span{Start: front[i].start, End: front[i].last.end}) {
			t.Fatal(ln)
		}
	}
}
