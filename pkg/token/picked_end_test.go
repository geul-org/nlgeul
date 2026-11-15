//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestPickedEnd — 지, the last letter of the closing 를 is written onto (pickedClosings), is such for 를; other letters are not, and no letter is for a form the table does not list
package token

import "testing"

func TestPickedEnd(t *testing.T) {
	if !pickedEnd("를", '지') {
		t.Fatal("지 for 를")
	}
	for _, r := range []rune{'고', '게', '다', '는', 'x', ':'} {
		if pickedEnd("를", r) {
			t.Errorf("%c for 를", r)
		}
	}
	if pickedEnd("의", '지') || pickedEnd("", '지') {
		t.Fatal("a form the table does not list")
	}
}
