//ff:func feature=token type=helper control=sequence
//ff:what TestSignsBefore — the run of name signs before a place; none when another letter stands there or at the start
package token

import "testing"

func TestSignsBefore(t *testing.T) {
	if signsBefore([]rune("(~/.a"), 4) != 1 || signsBefore([]rune("(a"), 1) != 1 || signsBefore([]rune("a"), 0) != 0 {
		t.Fatal(signsBefore([]rune("(~/.a"), 4))
	}
}
