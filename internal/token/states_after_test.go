//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestStatesAfter — the states at a span that have a state of the set before them (all of those at the first span), in the order given; the states at other spans are not
package token

import "testing"

func TestStatesAfter(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	pt := r.sc.between(0, 8)
	first := statesAfter(pt.states, Span{0, 4}, nil, true)
	next := statesAfter(pt.states, Span{4, 8}, first, false)
	if len(first) == 0 || len(next) == 0 {
		t.Fatal(len(first), len(next))
	}
	for _, s := range next {
		if s.last.end != 8 {
			t.Fatal(s.last.end)
		}
	}
	if got := statesAfter(pt.states, Span{4, 8}, first[:0], false); len(got) != 0 {
		t.Fatal("no state before", got)
	}
}
