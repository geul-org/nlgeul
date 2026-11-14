//ff:func feature=token type=reader control=sequence
//ff:what TestScanStretchOf — a portion of the net as a stretch: its letters, one cell for each span in the order of the spans, the candidates in the order of lambda and then of birth, the pairs of spans, the number of ways and the keys counted; only at the place where the unit begins (the first portion, a span that starts at 0) is a sign of a pair given once for each way it is read
package token

import "testing"

func TestScanStretchOf(t *testing.T) {
	sc := &scan{ix: scanOf("x").ix, s: []rune("가나")}
	root := sc.registry().root
	head := func(end int, p float64) step {
		e := wordStep(end, "(", 1, p)
		e.word.Parts[0].Name.Kind = pairName
		return e
	}
	a := addState(sc, []*path{root}, head(1, -1))
	b := addState(sc, []*path{a}, head(2, -2))
	pt := portion{from: 0, to: 2, states: []*path{a, b}}
	u := unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 1}, {openName, 2}}}
	st := sc.stretchOf(u, pt, true)
	if st.contest.Text != "가나" || st.cells == nil || len(st.cells.nodes) != 2 || len(st.cells.nodes[0].branches) != 2 || len(st.cells.nodes[1].branches) != 1 {
		t.Fatal("spread at the unit's head only", st.cells.nodes)
	}
	if st.cells.nodes[1].branches[0].Words[0].Parts[0].Name.Kind != pairName || st.cells.ways.Int64() != 1 || len(st.cells.edges) != 1 || st.cells.sets == 0 || st.cells.from != 0 || st.cells.clashes != 0 {
		t.Fatal(st.cells)
	}
	later := sc.stretchOf(u, pt, false)
	if len(later.cells.nodes[0].branches) != 1 {
		t.Fatal("not the first portion", later.cells.nodes)
	}
	moved := sc.stretchOf(u, portion{from: 1, to: 2, states: []*path{b}}, true)
	if moved.contest.Text != "나" || moved.cells.from != 1 || len(moved.cells.nodes[0].branches) != 1 {
		t.Fatal("a span that does not start at 0", moved.cells.nodes)
	}
}
