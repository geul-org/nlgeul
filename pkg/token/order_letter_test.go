//ff:func feature=token type=reader control=sequence
//ff:what TestOrderLetter — a letter of orderLetters right before a place is a number, taken as one of the 28 letters, at a word's head, after a sign and right after a number's end (a unit after a digit, a digit); after a syllable or a Latin letter, at the head, or a letter not of the list it is not
package token

import (
	"math"
	"testing"
)

func TestOrderLetter(t *testing.T) {
	pref := math.Log10(1.0 / 28)
	if got := scanOf("가목").orderLetter(1); len(got) != 1 || got[0].text != "가" || math.Abs(got[0].pref-pref) > 1e-12 {
		t.Fatal(got)
	}
	if got := scanOf("허목").orderLetter(1); len(got) != 1 || got[0].text != "허" {
		t.Fatal("the last letter of the list", got)
	}
	if got := scanOf("(나목").orderLetter(2); len(got) != 1 || got[0].text != "나" {
		t.Fatal("after a sign", got)
	}
	if got := scanOf("제1호가목").orderLetter(4); len(got) != 1 || got[0].text != "가" {
		t.Fatal("after a unit after a digit", got)
	}
	if got := scanOf("3가목").orderLetter(2); len(got) != 1 {
		t.Fatal("after a digit", got)
	}
	if got := scanOf("이가목").orderLetter(2); got != nil {
		t.Fatal("after a syllable", got)
	}
	if got := scanOf("A가목").orderLetter(2); got != nil {
		t.Fatal("after a Latin letter", got)
	}
	if got := scanOf("목").orderLetter(0); got != nil {
		t.Fatal("at the head", got)
	}
	if got := scanOf("곰목").orderLetter(1); got != nil {
		t.Fatal("not of the list", got)
	}
}
