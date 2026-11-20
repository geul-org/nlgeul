//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestLastOfAnyPicked — 지, the last letter of the closing 를 is written onto (pickedClosings), is such for 를; other letters are not, and no letter is for a form the table does not list
package parse

import "testing"

func TestLastOfAnyPicked(t *testing.T) {
	if !lastOfAny(pickedClosings["를"], '지') {
		t.Fatal("지 for 를")
	}
	for _, r := range []rune{'고', '게', '다', '는', 'x', ':'} {
		if lastOfAny(pickedClosings["를"], r) {
			t.Errorf("%c for 를", r)
		}
	}
	if lastOfAny(pickedClosings["의"], '지') || lastOfAny(pickedClosings[""], '지') {
		t.Fatal("a form the table does not list")
	}
}
