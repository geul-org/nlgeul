//ff:func feature=token type=reader control=sequence
//ff:what TestFitsMyriad — article 44: a number that ends in 만, 억 or 조 before a word that begins with a number fits; one that ends in another letter, or before a word that is no number, does not
package token

import "testing"

func TestFitsMyriad(t *testing.T) {
	if !fitsMyriad(gap{left: spacingWord("3/수", "억/큰 수:억"), right: spacingWord("5/수", "천/큰 수:천", "만/큰 수:만", "원/단위:원")}) || !fitsMyriad(gap{left: spacingWord("삼억오천만/수"), right: spacingWord("이천/수")}) {
		t.Fatal("after 억, after 만")
	}
	if fitsMyriad(gap{left: spacingWord("삼십/수"), right: spacingWord("5/수")}) || fitsMyriad(gap{left: spacingWord("3/수", "억/큰 수:억"), right: spacingWord("원/열쇠:체언")}) || fitsMyriad(gap{left: spacingWord("씨/열쇠:체언"), right: spacingWord("5/수")}) {
		t.Fatal("not after 만 억 조, not before a number, no number")
	}
}
