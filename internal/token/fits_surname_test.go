//ff:func feature=token type=reader control=sequence
//ff:what TestFitsSurname — the proviso of article 48: the word before the boundary is a family name of the list with no part after it; a family name with a particle, another stem or no stem does not
package token

import "testing"

func TestFitsSurname(t *testing.T) {
	if !fitsSurname(gap{left: spacingWord("김/열쇠:체언")}) || !fitsSurname(gap{left: spacingWord("황보")}) {
		t.Fatal("a family name")
	}
	if fitsSurname(gap{left: spacingWord("김/열쇠:체언", "이/주격:이")}) || fitsSurname(gap{left: spacingWord("사과/열쇠:체언")}) || fitsSurname(gap{left: spacingWord("-김/주격:가")}) {
		t.Fatal("a particle, another stem, no stem")
	}
}
