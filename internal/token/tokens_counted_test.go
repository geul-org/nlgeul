//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestTokensCounted — the tokens of a stem from b to end are the same for a scan that was prepared (from its counts) and one that was not (counted place by place): an atom counts as one; none for a stem of no length
package token

import "testing"

func TestTokensCounted(t *testing.T) {
	text := []rune("abc가나다def라")
	ready := scanOf(string(text))
	bare := &scan{s: text, inside: ready.inside}
	for b := 0; b <= len(text); b++ {
		for e := 0; e <= len(text); e++ {
			if ready.tokens(b, e) != bare.tokens(b, e) {
				t.Fatal(b, e, ready.tokens(b, e), bare.tokens(b, e))
			}
		}
	}
	if ready.tokens(0, len(text)) != 6 || ready.tokens(3, 3) != 0 || ready.tokens(5, 2) != 0 {
		t.Fatal(ready.tokens(0, len(text)))
	}
}
