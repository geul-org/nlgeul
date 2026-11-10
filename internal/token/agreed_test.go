//ff:func feature=token type=reader control=sequence
//ff:what TestAgreed — only the ends every reading has, in order
package token

import "testing"

func TestAgreed(t *testing.T) {
	a := grow(grow(grow(&path{verdict: 1}, step{end: 2}), step{end: 3}), step{end: 5})
	b := grow(grow(&path{verdict: 1}, step{end: 2}), step{end: 5})
	if got := agreed([]*path{a, b}); len(got) != 2 || got[0] != 2 || got[1] != 5 {
		t.Fatal(got)
	}
}
