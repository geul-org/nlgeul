//ff:func feature=token type=helper control=sequence
//ff:what TestPooled — a candidate printed the same as one before it is left out (the first stays in its place); a word printed the same with another verdict or preference is counted once however many there are, one with the same values is not counted
package token

import "testing"

func TestPooled(t *testing.T) {
	mk := func(text string, v, p float64) Branch {
		return Branch{Words: []Eojeol{{Parts: []Part{{Text: text, Stem: true}}}}, Verdict: v, Pref: p}
	}
	got, clashes := pooled([]Branch{mk("a", 1, -1), mk("b", 1, -2), mk("a", 1, -1), mk("c", 1, -3)})
	if len(got) != 3 || got[0].String() != "a" || got[1].String() != "b" || got[2].String() != "c" || clashes != 0 {
		t.Fatal(got, clashes)
	}
	got, clashes = pooled([]Branch{mk("a", 1, -1), mk("a", 0.5, -1), mk("a", 1, -2), mk("b", 1, -2), mk("b", 1, -3)})
	if len(got) != 2 || got[0].Pref != -1 || got[0].Verdict != 1 || clashes != 2 {
		t.Fatal(got, clashes)
	}
}
