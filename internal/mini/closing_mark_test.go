//ff:func feature=mini type=reader control=sequence
//ff:what TestClosingMark — a one-piece closing, a closing of several pieces, and the joining 으 as a piece of its own
package mini

import "testing"

func TestClosingMark(t *testing.T) {
	if got := closingMark(ending{form: "지만"}, "지만"); got != "+지만" {
		t.Fatal(got)
	}
	if got := closingMark(ending{form: "다고"}, "다고"); got != "+다+고" {
		t.Fatal(got)
	}
	if got := closingMark(ending{form: "면", flags: fEu}, "으면"); got != "+으+면" {
		t.Fatal(got)
	}
}
