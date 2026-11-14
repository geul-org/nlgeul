//ff:func feature=token type=helper control=sequence
//ff:what TestBracketOuter — "┐" at the first line, "┘" at the last, "│" between, the middle line (halved, rounded down) "├" with the tag; with three ways the first line of the one between is "┤", or "┼" when that is the middle line; one way alone is a single "├"
package token

import "testing"

func TestBracketOuter(t *testing.T) {
	two := newBracket(8)
	two.drawOuter([]int{1, 5}, "T")
	if string(two.outer) != " ┐│├│┘  " || two.tags[3] != "T" {
		t.Fatalf("%q %q", string(two.outer), two.tags)
	}
	three := newBracket(9)
	three.drawOuter([]int{0, 4, 8}, "T")
	if string(three.outer) != "┐│││┼│││┘" || three.tags[4] != "T" {
		t.Fatalf("%q %q", string(three.outer), three.tags)
	}
	four := newBracket(12)
	four.drawOuter([]int{0, 3, 7, 11}, "T")
	if string(four.outer) != "┐││┤│├│┤│││┘" || four.tags[5] != "T" {
		t.Fatalf("%q %q", string(four.outer), four.tags)
	}
	one := newBracket(4)
	one.drawOuter([]int{2}, "T")
	if string(one.outer) != "  ├ " || one.tags[2] != "T" {
		t.Fatalf("%q %q", string(one.outer), one.tags)
	}
}
