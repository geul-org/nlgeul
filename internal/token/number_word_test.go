//ff:func feature=token type=helper control=sequence
//ff:what TestNumberWord — a number word: a number stem with no part after it, or greater numbers and approximations only; not one with a unit or a particle after it, nor a stem that is no number
package token

import "testing"

func TestNumberWord(t *testing.T) {
	if !numberWord(spacingWord("세/수")) || !numberWord(spacingWord("3/수", "억/큰 수:억")) || !numberWord(spacingWord("10/수", "여/어림:여")) {
		t.Fatal("a number")
	}
	if numberWord(spacingWord("3/수", "개/단위:개")) || numberWord(spacingWord("3/수", "를/목적격:를")) || numberWord(spacingWord("씨/열쇠:체언")) {
		t.Fatal("not a number")
	}
}
