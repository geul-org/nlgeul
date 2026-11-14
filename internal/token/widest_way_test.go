//ff:func feature=token type=helper control=sequence
//ff:what TestWidestWay — the widest width inside over the boxes of all the ways; none give 0
package token

import "testing"

func TestWidestWay(t *testing.T) {
	a := boxWay{boxes: []box{{width: 20}, {width: 31}}}
	b := boxWay{boxes: []box{{width: 40}}}
	if got := widestWay([]boxWay{a, b}); got != 40 {
		t.Fatal(got)
	}
	if got := widestWay([]boxWay{b, a}); got != 40 {
		t.Fatal(got)
	}
	if got := widestWay([]boxWay{a}); got != 31 {
		t.Fatal(got)
	}
	if got := widestWay(nil); got != 0 {
		t.Fatal(got)
	}
}
