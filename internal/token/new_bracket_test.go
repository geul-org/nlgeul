//ff:func feature=token type=helper control=sequence
//ff:what TestNewBracket — a blank for each of the lines, in both brackets, and no tag
package token

import "testing"

func TestNewBracket(t *testing.T) {
	b := newBracket(3)
	if string(b.inner) != "   " || string(b.outer) != "   " || len(b.tags) != 3 || b.tags[0] != "" {
		t.Fatal(b)
	}
	if got := newBracket(0); len(got.inner) != 0 {
		t.Fatal(got)
	}
}
