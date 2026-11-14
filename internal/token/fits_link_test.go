//ff:func feature=token type=reader control=sequence
//ff:what TestFitsLink — article 45: the word after the boundary has a linking or closing word of the list for its stem, or the word before it is such a word with no part after it; no other word fits
package token

import "testing"

func TestFitsLink(t *testing.T) {
	if !fitsLink(gap{left: spacingWord("뚜루뚜"), right: spacingWord("겸/열쇠:체언")}) || !fitsLink(gap{left: spacingWord("뚜루뚜"), right: spacingWord("등/열쇠:체언", "을/목적격:을")}) {
		t.Fatal("before the word")
	}
	if !fitsLink(gap{left: spacingWord("및/열쇠:부사"), right: spacingWord("운영/열쇠:체언")}) {
		t.Fatal("after the word")
	}
	if fitsLink(gap{left: spacingWord("및/열쇠:부사", "도/보조사:도"), right: spacingWord("운영/열쇠:체언")}) || fitsLink(gap{left: spacingWord("뚜루뚜"), right: spacingWord("운영/열쇠:체언")}) || fitsLink(gap{left: spacingWord("-등/주격:가"), right: spacingWord("운영/열쇠:체언")}) {
		t.Fatal("a part after, no word of the list, no stem")
	}
}
