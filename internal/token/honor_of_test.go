//ff:func feature=token type=reader control=sequence
//ff:what TestHonorOf — 셨 is always the honorific's and asks nothing of the word before; 신 is read only with a word table, after a predicate's stem; another letter is not read
package token

import "testing"

func TestHonorOf(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if e, gate := ix.honorOf('셨'); e != ix.honor || gate != nil {
		t.Fatal(e)
	}
	if e, gate := ix.honorOf('신'); e == nil || e != ix.wide['신'] || gate != ix.pred {
		t.Fatal(e)
	}
	if e, _ := ix.honorOf('곰'); e != nil {
		t.Fatal(e)
	}
	if e, _ := scanOf("x").ix.honorOf('신'); e != nil {
		t.Fatal("no word table", e)
	}
}
