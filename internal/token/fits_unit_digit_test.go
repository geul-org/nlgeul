//ff:func feature=token type=reader control=sequence
//ff:what TestFitsUnitDigit — the proviso of article 43: a number that begins with a digit before a stem that is a unit fits; a Hangul number, or a stem that is no unit, does not
package token

import "testing"

func TestFitsUnitDigit(t *testing.T) {
	if !fitsUnitDigit(gap{left: spacingWord("3/수"), right: spacingWord("개/열쇠:체언", "를/목적격:를")}) {
		t.Fatal("3 개")
	}
	if fitsUnitDigit(gap{left: spacingWord("세/수"), right: spacingWord("마리/열쇠:체언")}) || fitsUnitDigit(gap{left: spacingWord("3/수"), right: spacingWord("사과/열쇠:체언")}) {
		t.Fatal("a Hangul number, no unit")
	}
}
