//ff:func feature=parse type=reader control=sequence
//ff:what TestStreamOuter — a span is given as it is when the stream has no table of places, and in the places with the spaces of a span of code left out when it has
package parse

import "testing"

func TestStreamOuter(t *testing.T) {
	plain := &Stream{}
	if got := plain.outer(Span{2, 5}); got != (Span{2, 5}) {
		t.Fatal(got)
	}
	st := &Stream{places: []int{0, 1, 2, 2, 3, 4}}
	if got := st.outer(Span{0, 5}); got != (Span{0, 4}) {
		t.Fatal(got)
	}
	if got := st.outer(Span{3, 4}); got != (Span{2, 3}) {
		t.Fatal(got)
	}
}
