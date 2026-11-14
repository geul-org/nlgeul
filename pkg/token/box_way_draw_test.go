//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBoxWayDraw — the boxes stacked at the width, every line width + 4 cells, last at the title line of the last box; drawn again it is drawn anew, and a way of one box ends at line 0
package token

import "testing"

func TestBoxWayDraw(t *testing.T) {
	seg := boxSegment(1, 2, 0, Span{0, 6}, boxNode(0, 3, "아버지", "아버지", "아버"), boxNode(3, 6, "가방에", "가방에"))
	w := newBoxWay(seg, 0)
	w.draw(40)
	if len(w.lines) != 4+3 || w.last != 4 {
		t.Fatal(len(w.lines), w.last)
	}
	for _, l := range w.lines {
		if cellWidth(l) != 44 {
			t.Fatalf("%q", l)
		}
	}
	w.draw(50)
	if len(w.lines) != 7 || w.last != 4 || cellWidth(w.lines[0]) != 54 {
		t.Fatal(len(w.lines), w.last)
	}
	one := newBoxWay(boxSegment(1, 2, 0, Span{0, 3}, boxNode(0, 3, "아버지", "아버지")), 0)
	one.draw(30)
	if one.last != 0 || len(one.lines) != 3 {
		t.Fatal(one.last, len(one.lines))
	}
}
