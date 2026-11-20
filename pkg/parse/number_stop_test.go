//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestNumberStop — a full stop right after a number is a piece of the number's word (31.까지, 1.5.에, 1.은 chain from the dot to the particle); no stop after a word, after a dot, in the decimal 1.5 or at the start; the particle chain reads it as a piece, so the stop is the particle's left neighbour
package parse

import "testing"

func TestNumberStop(t *testing.T) {
	sc := scanOf("31.까지")
	got := sc.numberStop(3)
	if len(got) != 1 || got[0].start != 2 || len(got[0].pieces) != 1 || got[0].pieces[0].text != "." || got[0].pieces[0].names != "마침표" || got[0].pieces[0].verdict <= 0 {
		t.Fatal(got)
	}
	for _, c := range []struct {
		text string
		j    int
		why  string
	}{
		{"했다.그는", 3, "after a word"},
		{"31..까지", 4, "after a dot"},
		{"1.5까지", 2, "the dot of a decimal is no stop"},
		{".가", 1, "at the start"},
		{"31.까지", 2, "no dot at j-1"},
	} {
		if g := scanOf(c.text).numberStop(c.j); len(g) != 0 {
			t.Fatal(c.why, g)
		}
	}
	var chain bool
	for _, c := range sc.chains(5, 0, -1) {
		chain = chain || (c.start == 2 && len(c.pieces) == 2 && c.pieces[0].text == "." && c.pieces[0].names == "마침표" && c.pieces[1].text == "까지")
	}
	if !chain {
		t.Fatal("no chain 31+.+까지")
	}
	var plain bool
	for _, c := range scanOf("했다.그는").chains(5, 0, -1) {
		plain = plain || (len(c.pieces) > 1 && c.pieces[0].text == ".")
	}
	if plain {
		t.Fatal("a stop after a word is joined to the particle")
	}
}
