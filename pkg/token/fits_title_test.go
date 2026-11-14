//ff:func feature=token type=reader control=sequence
//ff:what TestFitsTitle — article 48: the word after the boundary has a title or an office of the list for its stem; no other stem fits
package token

import "testing"

func TestFitsTitle(t *testing.T) {
	if !fitsTitle(gap{right: spacingWord("씨/열쇠:체언", "가/주격:가")}) || !fitsTitle(gap{right: spacingWord("총무/열쇠:체언")}) {
		t.Fatal("a title, an office")
	}
	if fitsTitle(gap{right: spacingWord("사과/열쇠:체언")}) {
		t.Fatal("no title")
	}
}
