//ff:func feature=token type=reader control=sequence
//ff:what TestFitsUnit — article 43: a number that does not begin with a digit before a stem that is a unit of the unit table fits; a digit, a unit after a word that is no number, or a stem that is no unit does not
package token

import "testing"

func TestFitsUnit(t *testing.T) {
	if !fitsUnit(gap{left: spacingWord("세/수"), right: spacingWord("마리/열쇠:체언")}) || !fitsUnit(gap{left: spacingWord("삼십/수"), right: spacingWord("개/열쇠:체언", "를/목적격:를")}) {
		t.Fatal("a Hangul number and a unit")
	}
	if fitsUnit(gap{left: spacingWord("3/수"), right: spacingWord("개/열쇠:체언")}) || fitsUnit(gap{left: spacingWord("씨/열쇠:체언"), right: spacingWord("개/열쇠:체언")}) || fitsUnit(gap{left: spacingWord("세/수"), right: spacingWord("사과/열쇠:체언")}) {
		t.Fatal("a digit, no number, no unit")
	}
}
