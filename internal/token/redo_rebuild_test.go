//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestRedoRebuild — a stretch a window cut off is read again from what the stream kept of it: the region has the ways the first reading had (the same words, in the same order), the first stretch from the root of the unit and a later one from the seeds it grew from, and with a sign at its end the candidates kept there are made
package token

import "testing"

func TestRedoRebuild(t *testing.T) {
	r, st := windowCut(t, 5)
	d := st.redo
	d.head, d.at = Contest{Text: st.contest.Text}, 0
	first := newRegionNet(d.head, 0, *st.cells)
	again := d.rebuild()
	a, _ := first.deal()
	b, _ := again.deal()
	if len(b.Nodes) != 1 || segmentLine(a) != segmentLine(b) || a.Lambda != b.Lambda || again.span != first.span {
		t.Fatal(segmentLine(a), segmentLine(b))
	}
	var second unitStretch
	for ok := false; !ok; {
		if second, ok = r.tryCut(); !ok {
			r.step()
		}
	}
	d = second.redo
	d.head, d.at = Contest{Text: second.contest.Text}, 4
	if len(d.seeds) == 0 || d.from != 4 {
		t.Fatal("grew from seeds", d.from)
	}
	a, _ = newRegionNet(d.head, 4, *second.cells).deal()
	b, _ = d.rebuild().deal()
	if segmentLine(a) != segmentLine(b) || a.Nodes[0].Span != (Span{4, 8}) {
		t.Fatal(segmentLine(a), segmentLine(b))
	}
	d.sc.mid[d.to] = midSign{}
	if c, _ := d.rebuild().deal(); segmentLine(c) != segmentLine(a) {
		t.Fatal("with the candidates kept at the sign")
	}
}
