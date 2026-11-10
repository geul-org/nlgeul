//ff:func feature=token type=helper control=sequence
//ff:what TestNested — brackets pair by nesting; one with no partner and one at a place to skip are left out
package token

import "testing"

func TestNested(t *testing.T) {
	r := []rune("((가)나) )(")
	got := nested(r, make([]bool, len(r)))
	if len(got) != 4 || got[0] != 5 || got[1] != 3 || got[3] != 1 {
		t.Fatal(got)
	}
	skip := make([]bool, len(r))
	skip[1] = true
	if got = nested(r, skip); got[0] != 3 {
		t.Fatal("a place to skip", got)
	}
}
