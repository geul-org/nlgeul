//ff:func feature=token type=helper control=sequence
//ff:what TestPartTokens — a run of Latin letters and digits is one token; a word of other letters has one for each
package token

import "testing"

func TestPartTokens(t *testing.T) {
	if got := partTokens("abc_12"); got != 1 {
		t.Fatal(got)
	}
	if got := partTokens("사자"); got != 2 {
		t.Fatal(got)
	}
	if got := partTokens("곰a"); got != 2 {
		t.Fatal(got)
	}
}
