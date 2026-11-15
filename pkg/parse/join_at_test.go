//ff:func feature=parse type=reader control=sequence
//ff:what TestJoinAt — a longer chain replaces the word's and tails are added; a shorter one changes nothing; a dependent noun of the table marked 붙여 읽음 is read on a modifier (먹은적이, 앉은채로, 할시에는), one marked - is not (하는간, 하는외; 올듯 and 웃을만도, whose 듯 and 만 the closing and particle tables read)
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
	w = &word{}
	joinAt(w, []rune("먹은적이"), 2)
	if w.chain.marked != "먹+은+적+이" || w.chain.size != 6 {
		t.Fatal(w.chain)
	}
	w = &word{}
	joinAt(w, []rune("앉은채로"), 2)
	if w.chain.marked != "앉+은+채+로" || w.chain.size != 6 {
		t.Fatal(w.chain)
	}
	w = &word{}
	joinAt(w, []rune("하는외"), 2)
	if w.chain.size != 0 {
		t.Fatal(w.chain)
	}
	w = &word{}
	joinAt(w, []rune("할시에는"), 1)
	if w.chain.marked != "할+시+에+는" || w.chain.size != 7 {
		t.Fatal(w.chain)
	}
	w = &word{}
	joinAt(w, []rune("올듯"), 1)
	if w.chain.size != 0 || len(w.tails) != 0 {
		t.Fatal(w.chain, w.tails)
	}
	w = &word{}
	joinAt(w, []rune("웃을만도"), 2)
	if w.chain.size != 0 || len(w.tails) != 0 {
		t.Fatal(w.chain, w.tails)
	}
}
