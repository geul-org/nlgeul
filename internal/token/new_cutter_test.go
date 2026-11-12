//ff:func feature=token type=reader control=sequence
//ff:what TestNewCutter — a cutter at a sentence's start holds the index, the letters, the spans of code (each backtick with its partner), the places they cover, one way of standing open with nothing open, and whether a trace is kept
package token

import "testing"

func TestNewCutter(t *testing.T) {
	ix, _ := newIndex(Options{})
	c := newCutter(ix, "가 `(` 나", true)
	if c.ix != ix || string(c.r) != "가 `(` 나" || !c.trace {
		t.Fatal("index, letters or trace", c)
	}
	if len(c.code) != 2 || c.code[2] != 4 || c.code[4] != 2 {
		t.Fatal("the span of code", c.code)
	}
	if len(c.inside) != len(c.r) || !c.inside[2] || !c.inside[3] || !c.inside[4] || c.inside[1] || c.inside[5] || c.inside[0] {
		t.Fatal("the places of the span of code, its backticks too", c.inside)
	}
	if len(c.open) != 1 || c.open[0] != "" {
		t.Fatal("one way, with nothing open", c.open)
	}
	plain := newCutter(ix, "", false)
	if plain.trace || len(plain.r) != 0 || len(plain.code) != 0 || len(plain.open) != 1 {
		t.Fatal("an empty sentence", plain)
	}
}
