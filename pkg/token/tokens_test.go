//ff:func feature=token type=reader control=sequence
//ff:what TestTokens — an atom counts as one token, Hangul syllables one each
package token

import "testing"

func TestTokens(t *testing.T) {
	sc := scanOf("task.py에서")
	if sc.tokens(0, 7) != 1 || sc.tokens(0, 9) != 3 || sc.tokens(7, 9) != 2 {
		t.Fatal(sc.tokens(0, 7), sc.tokens(0, 9))
	}
}
