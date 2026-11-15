//ff:func feature=rule type=helper control=sequence
//ff:what TestNearAt — the i-th word on one side is given (0 the nearest); past the last one, and on a side with no word, an empty one
package rule

import "testing"

func TestNearAt(t *testing.T) {
	ns := []Near{{Stem: "가"}, {Stem: "나", Pred: true}}
	if nearAt(ns, 0).Stem != "가" || nearAt(ns, 1) != (Near{Stem: "나", Pred: true}) {
		t.Fatal(nearAt(ns, 0), nearAt(ns, 1))
	}
	if nearAt(ns, 2) != (Near{}) || nearAt(nil, 0) != (Near{}) {
		t.Fatal("none so far")
	}
}
