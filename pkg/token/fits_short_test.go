//ff:func feature=token type=reader control=sequence
//ff:what TestFitsShort — article 46: both words are one letter long without their particles fits; a word of two letters or two letters left after a particle does not
package token

import "testing"

func TestFitsShort(t *testing.T) {
	if !fitsShort(gap{left: spacingWord("개/열쇠:체언"), leftLen: 1, right: spacingWord("선/열쇠:체언"), rightLen: 1}) {
		t.Fatal("개 선")
	}
	if !fitsShort(gap{left: spacingWord("씨/열쇠:체언", "가/주격:가"), leftLen: 2, right: spacingWord("선/열쇠:체언"), rightLen: 1}) {
		t.Fatal("a particle is not counted")
	}
	if fitsShort(gap{left: spacingWord("가방/열쇠:체언"), leftLen: 2, right: spacingWord("선/열쇠:체언"), rightLen: 1}) || fitsShort(gap{left: spacingWord("개/열쇠:체언"), leftLen: 1, right: spacingWord("선/열쇠:체언", "에서/부사격:에서"), rightLen: 4}) {
		t.Fatal("two letters")
	}
}
