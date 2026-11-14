//ff:func feature=token type=helper control=sequence
//ff:what TestBracketInner — "┬" at the first line, "│" between, "┘" at the last, the tag on the first; a way of one line is "─"; the lines around are left blank
package token

import "testing"

func TestBracketInner(t *testing.T) {
	b := newBracket(6)
	b.drawInner(1, 4, "T")
	if string(b.inner) != " ┬││┘ " || b.tags[1] != "T" || b.tags[4] != "" {
		t.Fatalf("%q %q", string(b.inner), b.tags)
	}
	c := newBracket(3)
	c.drawInner(1, 1, "U")
	if string(c.inner) != " ─ " || c.tags[1] != "U" {
		t.Fatalf("%q %q", string(c.inner), c.tags)
	}
	d := newBracket(3)
	d.drawInner(0, 1, "V")
	if string(d.inner) != "┬┘ " {
		t.Fatalf("%q", string(d.inner))
	}
}
