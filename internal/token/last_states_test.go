//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestLastStates — the states a stretch's best way ends in are those at its last span that a line of the spans leads to; none for a way of no spans
package token

import "testing"

func TestLastStates(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	pt := r.sc.between(0, 8)
	seg := Segment{Nodes: []Node{{Span: Span{0, 4}}, {Span: Span{4, 8}}}}
	got := lastStates(pt, seg)
	if len(got) == 0 {
		t.Fatal("none")
	}
	for _, s := range got {
		if s.start != 4 || s.last.end != 8 {
			t.Fatal(s.start, s.last.end)
		}
	}
	if got = lastStates(pt, Segment{}); got != nil {
		t.Fatal(got)
	}
}
