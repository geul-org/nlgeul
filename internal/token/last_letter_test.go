//ff:func feature=token type=helper control=sequence
//ff:what TestLastLetter — the last letter of the last part of a word; 0 for a word with no part and for a last part with no letters (a dropped 이)
package token

import "testing"

func TestLastLetter(t *testing.T) {
	if got := lastLetter(spacingWord("삼/수", "억/큰 수:억")); got != '억' {
		t.Fatal(got)
	}
	if got := lastLetter(Eojeol{}); got != 0 {
		t.Fatal(got)
	}
	if got := lastLetter(spacingWord("총무/열쇠:체언", "/서술격:이(빠짐)")); got != 0 {
		t.Fatal(got)
	}
}
