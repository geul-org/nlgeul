//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunScore — a state takes the best way among the states before it that have one (a root is worth 0, a state with no way yet is passed over, the first to come wins a tie), and adds the best lambda of its span
package token

import "testing"

func TestNetRunScore(t *testing.T) {
	r, _ := windowRun(t, "words-far.tsv", "고양이")
	sc := r.sc
	root := sc.registry().root
	a := addState(sc, []*path{root}, wordStep(1, "a", 1, -3))
	b := addState(sc, []*path{root}, wordStep(1, "b", 1, -1))
	c := addState(sc, []*path{root}, wordStep(1, "c", 1, -1))
	lost := &path{start: 0, last: wordStep(1, "lost", 1, 0), preds: []*path{root}, prev: root}
	s := addState(sc, []*path{a, b, c}, wordStep(3, "s", 1, -2))
	s.preds = append([]*path{lost}, s.preds...)
	rc := &reach{value: map[*path]float64{}, back: map[*path]*path{}, group: map[Span][]*path{}}
	for _, p := range []*path{a, b, c} {
		rc.group[Span{Start: 0, End: 1}] = append(rc.group[Span{Start: 0, End: 1}], p)
		r.score(rc, p)
	}
	rc.group[Span{Start: 1, End: 3}] = []*path{s}
	r.score(rc, s)
	if rc.back[a] != root || rc.back[s] != b {
		t.Fatal("the best of the states before, the first on a tie", rc.back[s])
	}
	if _, ok := rc.value[lost]; ok {
		t.Fatal("no value for a state passed over")
	}
}
