//ff:func feature=mini type=rule control=sequence
//ff:what TestNoteSplit — a chain claim notes the particle reading; a tail claim notes its tail's division, the longest tail winning
package mini

import "testing"

func TestNoteSplit(t *testing.T) {
	r := &ruling{}
	noteSplit(r, "closingPredicate", tailFact{size: 3, marked: "한+다"})
	noteSplit(r, "closingPredicate", tailFact{size: 2, marked: "한다"})
	if r.chain || r.tail != "한+다" {
		t.Fatal(r)
	}
	noteSplit(r, "subjectCase", tailFact{size: 9, marked: "x"})
	if !r.chain || r.tail != "한+다" {
		t.Fatal(r)
	}
}
