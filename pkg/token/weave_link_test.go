//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveLink — the states before each state that lie in the portion (a state before that is not in it is left), and for each span the spans that follow it, each once and the nearest first
package token

import "testing"

func TestWeaveLink(t *testing.T) {
	w, s := oddWeave()
	w.link(map[*path]int{s["c"]: 0, s["x"]: 1, s["z"]: 2, s["y"]: 3, s["q"]: 4, s["t"]: 5})
	if len(w.preds[0]) != 0 || len(w.preds[2]) != 1 || w.preds[2][0] != 0 {
		t.Fatal(w.preds)
	}
	if len(w.succ[0]) != 1 || w.succ[0][0] != 2 || len(w.succ[1]) != 1 || w.succ[1][0] != 4 || len(w.succ[5]) != 0 {
		t.Fatal(w.succ)
	}
	w.link(map[*path]int{s["z"]: 2, s["y"]: 3})
	if len(w.preds[3]) != 1 || len(w.preds[2]) != 0 || len(w.succ[0]) != 0 {
		t.Fatal("a state before that is not in the portion", w.preds, w.succ)
	}
	dup, ds := oddWeave()
	extra := addState(&scan{}, []*path{ds["c"]}, wordStep(2, "z", 1, -1))
	extra.start = 1
	dup.states[2].preds = append(dup.states[2].preds, extra)
	dup.link(map[*path]int{ds["c"]: 0, extra: 0, ds["z"]: 2})
	if len(dup.succ[0]) != 1 {
		t.Fatal("each pair of spans once", dup.succ)
	}
}
