//ff:func feature=token type=reader control=sequence
//ff:what TestFarNew — a stem is too long to be a new word when it holds more than newStemMax tokens: eleven letters are, ten are not, an atom counts as one token however long, the whole unit is not (sole), the word at the unit's head (a sign of a pair or a span of code) is not
package token

import "testing"

func TestFarNew(t *testing.T) {
	sc := scanOf("가나다라마바사아자차카를")
	if !sc.farNew(0, tail{start: 11}, false) || sc.farNew(1, tail{start: 11}, false) || sc.farNew(0, tail{start: 10}, false) {
		t.Fatal("eleven letters, ten and ten")
	}
	if sc.farNew(0, tail{start: 11}, true) {
		t.Fatal("the whole unit")
	}
	head := scanOf("가나다라마바사아자차카를")
	head.head = 11
	if head.farNew(0, tail{start: 11}, false) || !head.farNew(0, tail{start: 12}, false) || !head.farNew(1, tail{start: 12}, false) && false {
		t.Fatal("the word at the unit's head, and a word that is not")
	}
	atom := scanOf("abcdefghijklmnopqrstuvwxyz가나다")
	if atom.farNew(0, tail{start: 26}, false) || atom.tokens(0, 29) != 4 {
		t.Fatal("an atom is one token", atom.tokens(0, 29))
	}
}
