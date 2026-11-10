//ff:func feature=token type=helper control=sequence
//ff:what TestWet — ㅏ ㅓ ㅑ ㅕ go to ㅐ ㅔ ㅒ ㅖ; a wide stem's syllable goes to ㅐ whatever its vowel
package token

import "testing"

func TestWet(t *testing.T) {
	if wet('라', false) != '래' || wet('러', false) != '레' || wet('야', false) != '얘' || wet('여', false) != '예' {
		t.Fatal(string(wet('라', false)), string(wet('러', false)))
	}
	if wet('러', true) != '래' || wet('떠', true) != '때' || wet('라', true) != '래' {
		t.Fatal(string(wet('러', true)))
	}
}
