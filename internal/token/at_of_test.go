//ff:func feature=token type=helper control=sequence
//ff:what TestAtOf — a letter of the text, and 0 before it and after it
package token

import "testing"

func TestAtOf(t *testing.T) {
	s := []rune("곰이")
	if atOf(s, 0) != '곰' || atOf(s, -1) != 0 || atOf(s, 2) != 0 {
		t.Fatal("atOf")
	}
}
