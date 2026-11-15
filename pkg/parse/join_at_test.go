//ff:func feature=parse type=reader control=sequence
//ff:what TestJoinAt — a longer chain replaces the word's and tails are added; a shorter one changes nothing; a dependent noun of the table marked - is not read on a modifier (하는간)
package parse

import "testing"

func TestJoinAt(t *testing.T) {
	w := &word{}
	joinAt(w, []rune("아는거"), 2)
	if w.chain.kase != caseBare || w.chain.size != 4 {
		t.Fatal(w.chain)
	}
	w = &word{chain: chainFact{size: 9}}
	joinAt(w, []rune("있는거지"), 2)
	if w.chain.size != 9 || len(w.tails) != 1 {
		t.Fatal(w.chain, w.tails)
	}
	w = &word{}
	joinAt(w, []rune("하는간"), 2)
	if w.chain.size != 0 {
		t.Fatal(w.chain)
	}
}
