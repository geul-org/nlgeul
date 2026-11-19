//ff:func feature=token type=rule control=sequence
//ff:what TestDropGate — before a quoting closing of droppableNew (라는, 라) the word before a dropped copula stem passes the gate of a nominal of the table or a new word (nominalNew), before the other closings of droppableKnown (네, 기) the gate of a nominal of the table (nominal)
package token

import "testing"

func TestDropGate(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	if ix.dropGate("꾸밈:라는") != ix.nominalNew || ix.dropGate("맺음:라") != ix.nominalNew || ix.nominalNew == nil {
		t.Fatal("the quoting closings")
	}
	if ix.dropGate("맺음:네") != ix.nominal || ix.dropGate("이름:기") != ix.nominal || ix.nominal == nil {
		t.Fatal("the other closings")
	}
}
