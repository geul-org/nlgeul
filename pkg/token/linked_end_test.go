//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestLinkedEnd — 서, the last letter of the linking 어서 (copulaClosings), is such; the last letters of other closings and other letters are not
package token

import "testing"

func TestLinkedEnd(t *testing.T) {
	if !linkedEnd('서') {
		t.Fatal("서")
	}
	for _, r := range []rune{'고', '면', '지', '어', '아', '다', 'x', ':'} {
		if linkedEnd(r) {
			t.Errorf("%c", r)
		}
	}
}
