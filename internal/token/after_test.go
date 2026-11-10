//ff:func feature=token type=helper control=sequence
//ff:what TestAfter — the runs that begin before a place are left out
package token

import "testing"

func TestAfter(t *testing.T) {
	sc := scanOf("x")
	if got := sc.after([][2]int{{0, 2}, {3, 5}, {6, 7}}, 5); len(got) != 1 || got[0][0] != 6 {
		t.Fatal(got)
	}
}
