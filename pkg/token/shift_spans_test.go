//ff:func feature=token type=helper control=sequence
//ff:what TestShiftSpans — each span moved by the same places, forward and back; none stays none
package token

import "testing"

func TestShiftSpans(t *testing.T) {
	got := shiftSpans([]Span{{0, 2}, {2, 5}}, 10)
	if len(got) != 2 || got[0] != (Span{10, 12}) || got[1] != (Span{12, 15}) {
		t.Fatal(got)
	}
	if back := shiftSpans(got, -10); back[0] != (Span{0, 2}) || back[1] != (Span{2, 5}) {
		t.Fatal(back)
	}
	if len(shiftSpans(nil, 3)) != 0 {
		t.Fatal("none")
	}
}
