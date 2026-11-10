//ff:func feature=token type=reader control=sequence
//ff:what TestNextUnit — the unit at a place by what stands there: a span of code, an opening sign, a closing sign with what follows, or plain letters up to the next sign of a pair
package token

import "testing"

func TestNextUnit(t *testing.T) {
	r := []rune("곰(`a`를)이")
	pairs := pairsOf(r)
	if u, end := nextUnit(r, 0, pairs); u.text != "곰" || u.head != 0 || end != 1 {
		t.Fatal(u, end)
	}
	if u, end := nextUnit(r, 1, pairs); u.text != "(" || u.name != "여는 묶음표" || end != 2 {
		t.Fatal(u, end)
	}
	if u, end := nextUnit(r, 2, pairs); u.text != "`a`를" || u.head != 3 || u.name != codeName || end != 6 {
		t.Fatal(u, end)
	}
	if u, end := nextUnit(r, 6, pairs); u.text != ")이" || u.head != 1 || u.name != "닫는 묶음표" || end != 8 {
		t.Fatal(u, end)
	}
}
