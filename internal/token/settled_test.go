//ff:func feature=token type=reader control=sequence
//ff:what TestSettled — with a word table, readings of known words only push out the others; without a known reading, or without a table, every reading stays
package token

import "testing"

func TestSettled(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	known, other := &path{verdict: 1}, &path{verdict: ix.unknown}
	if got := ix.settled([]*path{known, other}); len(got) != 1 || got[0] != known {
		t.Fatal(got)
	}
	if got := ix.settled([]*path{other, other}); len(got) != 2 {
		t.Fatal(got)
	}
	if got := ix.settled(nil); got != nil {
		t.Fatal(got)
	}
	if got := scanOf("x").ix.settled([]*path{known, other}); len(got) != 2 {
		t.Fatal(got)
	}
}
