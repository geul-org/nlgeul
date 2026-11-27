//ff:func feature=parse type=reader control=sequence
//ff:what TestBigFront — a greater number written in Hangul right before digits that a greater number follows is read as that greater number (천500억, 천6백); not where a unit follows the digits (만19세, 조1위), a numeral or a number stands before the letter (일천500억, 3천500억), a piece is written onto it, or the letter is no greater number
package parse

import "testing"

func TestBigFront(t *testing.T) {
	sc := scanOf("천500억")
	got := sc.bigFront(0, "천", tail{start: 1})
	if len(got) != 1 || got[0].name != "큰 수:천" || got[0].v != 1 || got[0].v != sc.ix.bigFront || got[0].kind != 0 {
		t.Fatal(got)
	}
	if got = scanOf("천6백여").bigFront(0, "천", tail{start: 1}); len(got) != 1 {
		t.Fatal(got)
	}
	if got = scanOf("만19세").bigFront(0, "만", tail{start: 1}); got != nil {
		t.Fatal("a unit after the digits", got)
	}
	if got = scanOf("조1위").bigFront(0, "조", tail{start: 1}); got != nil {
		t.Fatal("a unit after the digits", got)
	}
	if got = scanOf("천500").bigFront(0, "천", tail{start: 1}); got != nil {
		t.Fatal("nothing after the digits", got)
	}
	if got = scanOf("일천500억").bigFront(1, "천", tail{start: 2}); got != nil {
		t.Fatal("a numeral before the letter", got)
	}
	if got = scanOf("3천500억").bigFront(1, "천", tail{start: 2}); got != nil {
		t.Fatal("a number before the letter", got)
	}
	if got = sc.bigFront(0, "천", tail{start: 1, pieces: []piece{{}}}); got != nil {
		t.Fatal("a piece on the letter", got)
	}
	if got = scanOf("곰500억").bigFront(0, "곰", tail{start: 1}); got != nil {
		t.Fatal("no greater number", got)
	}
	if got = scanOf("천만에").bigFront(0, "천", tail{start: 1}); got != nil {
		t.Fatal("no digits after the letter", got)
	}
}
