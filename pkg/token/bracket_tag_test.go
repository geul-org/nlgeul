//ff:func feature=token type=helper control=sequence
//ff:what TestBracketTag — a tag is put on its line; a second one on the same line follows the first with two spaces between
package token

import "testing"

func TestBracketTag(t *testing.T) {
	b := newBracket(2)
	b.tag(1, "a")
	if b.tags[0] != "" || b.tags[1] != "a" {
		t.Fatal(b.tags)
	}
	b.tag(1, "b")
	if b.tags[1] != "a  b" {
		t.Fatal(b.tags)
	}
}
