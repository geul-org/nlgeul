//ff:func feature=token type=helper control=sequence
//ff:what TestStretchEnd — letters stop at a space, at a sign of a pair, or at the end
package token

import "testing"

func TestStretchEnd(t *testing.T) {
	r := []rune("가나 다(라")
	if stretchEnd(r, 0, nil) != 2 || stretchEnd(r, 3, map[int]int{4: 9}) != 4 || stretchEnd(r, 3, nil) != 6 {
		t.Fatal(stretchEnd(r, 0, nil), stretchEnd(r, 3, nil))
	}
}
