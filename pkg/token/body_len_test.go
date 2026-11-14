//ff:func feature=token type=helper control=sequence
//ff:what TestBodyLen — the letters a word takes up less those of its particle parts; the copula and a closing are not taken off
package token

import "testing"

func TestBodyLen(t *testing.T) {
	if got := bodyLen(spacingWord("씨/열쇠:체언", "가/주격:가"), 2); got != 1 {
		t.Fatal(got)
	}
	if got := bodyLen(spacingWord("가방/열쇠:체언"), 2); got != 2 {
		t.Fatal(got)
	}
	if got := bodyLen(spacingWord("씨/열쇠:체언", "이/서술격:이", "다/맺음:다"), 3); got != 3 {
		t.Fatal(got)
	}
}
