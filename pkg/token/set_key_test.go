//ff:func feature=token type=helper control=sequence
//ff:what TestSetKey — a span and a set of states as a key: the same for the same, another for another span, another set, another order, and one that does not run two numbers together (a number over 127 takes two bytes)
package token

import "testing"

func TestSetKey(t *testing.T) {
	a := setKey(3, []int{1, 2})
	if a != setKey(3, []int{1, 2}) || a == setKey(4, []int{1, 2}) || a == setKey(3, []int{1, 3}) || a == setKey(3, []int{2, 1}) || a == setKey(3, []int{1}) {
		t.Fatal("told apart")
	}
	if setKey(0, []int{300}) == setKey(0, []int{44, 2}) || len(setKey(0, []int{300})) != 3 {
		t.Fatal("one number each")
	}
}
