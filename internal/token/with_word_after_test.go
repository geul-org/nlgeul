//ff:func feature=token type=helper control=sequence
//ff:what TestWithWordAfter — the common spans stay when every line goes on after them, lose the last when a line ends with it (and so on to none for a line of one span or none)
package token

import "testing"

func TestWithWordAfter(t *testing.T) {
	a, b, c := Span{0, 4}, Span{4, 8}, Span{8, 12}
	if got := withWordAfter([]Span{a, b}, [][]Span{{a, b, c}, {a, b, c}}); len(got) != 2 {
		t.Fatal("goes on", got)
	}
	if got := withWordAfter([]Span{a, b}, [][]Span{{a, b, c}, {a, b}}); len(got) != 1 || got[0] != a {
		t.Fatal("a line ends with the last", got)
	}
	if got := withWordAfter([]Span{a}, [][]Span{{a}}); len(got) != 0 {
		t.Fatal("a line of one span", got)
	}
	if got := withWordAfter(nil, [][]Span{{}}); len(got) != 0 {
		t.Fatal("a line of none", got)
	}
}
