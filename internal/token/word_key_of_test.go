//ff:func feature=token type=helper control=sequence
//ff:what TestWordKeyOf — a candidate's last word by where it begins (where the one before ends), where it ends and how it is printed; the root gives place 0 as where the first word begins
package token

import "testing"

func TestWordKeyOf(t *testing.T) {
	p := grow(grow(&path{verdict: 1}, wordStep(2, "a", 1, -1)), wordStep(5, "b", 1, -1))
	if got := wordKeyOf(p); got != (wordKey{start: 2, end: 5, sig: "b"}) {
		t.Fatal(got)
	}
	if got := wordKeyOf(p.prev); got != (wordKey{start: 0, end: 2, sig: "a"}) {
		t.Fatal(got)
	}
}
