//ff:func feature=parse type=helper control=sequence
//ff:what TestFirstRune — the first letter of a text, whatever its width; the replacement character for an empty text
package parse

import "testing"

func TestFirstRune(t *testing.T) {
	if firstRune("kg") != 'k' || firstRune("ｍ") != 'ｍ' || firstRune("") != '�' {
		t.Fatal("firstRune")
	}
}
