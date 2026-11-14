//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveBranches — the candidates at a span, one for each state in the order made, each one word with its verdict, preference and number of birth; two states that print the same word with other values leave the better by the order candidates have (the earlier when equal), in the place of the first
package token

import "testing"

func TestWeaveBranches(t *testing.T) {
	w, _ := oddWeave()
	got := w.branches(w.at[Span{0, 2}])
	if len(got) != 1 || got[0].String() != "x" || got[0].Verdict != 1 || got[0].Pref != -2 || len(got[0].scores) != 1 {
		t.Fatal(got)
	}
	sc := &scan{ix: scanOf("x").ix}
	root := sc.registry().root
	worse := addState(sc, []*path{root}, wordStep(2, "same", 1, -5))
	other := addState(sc, []*path{root}, wordStep(2, "other", 1, -4))
	better := addState(sc, []*path{root}, wordStep(2, "same", 1, -3))
	later := addState(sc, []*path{root}, wordStep(2, "same", 1, -4))
	pt := portion{from: 0, to: 2, states: []*path{worse, other, better, later}}
	tw := newWeave(sc.ix, pt)
	got = tw.branches(0)
	if len(got) != 2 || got[0].String() != "same" || got[0].Pref != -3 || got[1].String() != "other" {
		t.Fatal(got)
	}
	tie := newWeave(sc.ix, portion{from: 0, to: 2, states: []*path{other, addState(sc, []*path{root}, wordStep(2, "other", 1, -4.0))}})
	if got = tie.branches(0); len(got) != 1 {
		t.Fatal("equal", got)
	}
}
