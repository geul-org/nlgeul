//ff:func feature=parse type=helper control=sequence
//ff:what TestDroppedIn — a tail with the copula's dropped stem among its pieces holds it, and a tail with the stem said, or with no copula, does not
package parse

import "testing"

func TestDroppedIn(t *testing.T) {
	da := piece{text: "다", names: "맺음:다"}
	if !droppedIn(tail{pieces: []piece{{text: "(이)", names: droppedName}, da}}) {
		t.Fatal("a dropped stem")
	}
	if droppedIn(tail{pieces: []piece{{text: "이", names: copulaName}, da}}) || droppedIn(tail{pieces: []piece{da}}) || droppedIn(tail{}) {
		t.Fatal("no dropped stem")
	}
}
