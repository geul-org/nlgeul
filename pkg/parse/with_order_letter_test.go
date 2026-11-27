//ff:func feature=parse type=helper control=sequence
//ff:what TestWithOrderLetter — the numbers before 목 with the order letter added: a numeral of the same letter is dropped (사 of 사목), a longer numeral and the others stay, and the letter is named the plain 수; with no letter the numbers are as they were
package parse

import "testing"

func TestWithOrderLetter(t *testing.T) {
	words := []numeral{{"삼사", -2, numName}, {"사", -1, numName}}
	letter := []numeral{{"사", -1.4, plainNumName}}
	got := withOrderLetter(words, letter)
	if len(got) != 2 || got[0].text != "삼사" || got[1].text != "사" || got[1].name != plainNumName {
		t.Fatal(got)
	}
	if got = withOrderLetter(words, nil); len(got) != 2 || got[1].name != numName {
		t.Fatal(got)
	}
}
