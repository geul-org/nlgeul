//ff:func feature=token type=helper control=sequence
//ff:what TestBracketText — a space, the two brackets, and a space and the tag when the line has one; blank brackets and no tag are three blanks
package token

import "testing"

func TestBracketText(t *testing.T) {
	b := newBracket(2)
	b.drawInner(0, 1, "① λ 합 -1.00")
	b.drawOuter([]int{0}, "x")
	if got := b.text(0); got != " ┬├ ① λ 합 -1.00  x" {
		t.Fatalf("%q", got)
	}
	if got := b.text(1); got != " ┘ " {
		t.Fatalf("%q", got)
	}
	if got := newBracket(1).text(0); got != "   " {
		t.Fatalf("%q", got)
	}
}
