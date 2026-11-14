//ff:func feature=token type=helper control=sequence
//ff:what TestNewBoxWay — a box for each node of the segment in order, not drawn yet; a segment with no node gives none
package token

import "testing"

func TestNewBoxWay(t *testing.T) {
	seg := boxSegment(1, 1, 0, Span{0, 6}, boxNode(0, 3, "아버지", "아버지"), boxNode(3, 6, "가방에", "가방", "가방에"))
	w := newBoxWay(seg, 0)
	if len(w.boxes) != 2 || w.boxes[1].title != "[3-6] 가방에 ─ 후보경합 2" || w.lines != nil {
		t.Fatal(w)
	}
	if got := newBoxWay(Segment{}, 0); len(got.boxes) != 0 {
		t.Fatal(got)
	}
}
