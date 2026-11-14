//ff:func feature=token type=reader control=sequence
//ff:what TestLetters — the letters from place b to place e are a part of the unit's string when the places were laid out (placeLetters), and made from the letters when they were not; both are the same text; the units of the tests are read by the net too
package token

import "testing"

func TestLetters(t *testing.T) {
	driveAll(t)
	sc := &scan{s: []rune("가a나")}
	if sc.letters(1, 3) != "a나" {
		t.Fatal("without places", sc.letters(1, 3))
	}
	sc.placeLetters()
	if sc.letters(1, 3) != "a나" || sc.letters(0, 3) != "가a나" || sc.letters(2, 2) != "" {
		t.Fatal("with places", sc.letters(1, 3))
	}
}
