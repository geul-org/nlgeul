//ff:func feature=token type=helper control=sequence
//ff:what TestFirstRune — the first letter of a text, whatever its width; the replacement character for an empty text
package token

import "testing"

func TestFirstRune(t *testing.T) {
	if firstRune("kg") != 'k' || firstRune("ｍ") != 'ｍ' || firstRune("") != '�' {
		t.Fatal("firstRune")
	}
}
