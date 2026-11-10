//ff:func feature=token type=helper control=sequence
//ff:what TestLettered — a place with a letter or digit on both sides; not at an edge or next to a space
package token

import "testing"

func TestLettered(t *testing.T) {
	r := []rune("a'b 'c")
	if !lettered(r, 1) || lettered(r, 4) || lettered(r, 0) || lettered(r, 5) {
		t.Fatal("lettered")
	}
}
