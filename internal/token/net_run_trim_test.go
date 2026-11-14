//ff:func feature=token type=helper control=sequence
//ff:what TestNetRunTrim — after a cut the places before it are emptied, the cut keeps the seeds, the places after it keep only the states that stay, the candidates kept at signs are put in order, the lambdas of the spans before it are forgotten and the scan is told where its first place is
package token

import "testing"

func TestNetRunTrim(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	r.sc.kept[2], r.sc.kept[6] = []*path{{}}, r.best[6]
	r.lam[Span{0, 4}], r.lam[Span{4, 8}] = spanLambda{}, spanLambda{}
	order := r.byEnd()
	seedOf, seeds := r.seedsOf(order, endingAt(order, 4))
	r.sc.kept[4] = r.best[4][:1]
	r.trim(4, seedOf, order, map[*path]bool{})
	if r.best[3] != nil || r.sc.marks[3] != nil || len(r.best[4]) != len(seeds) || len(r.best[8]) != 0 || r.sc.floor != 4 || r.seeds[0] != seeds[0] {
		t.Fatal(r.best[3], r.best[4], r.sc.floor)
	}
	if _, ok := r.sc.kept[2]; ok || len(r.sc.kept[6]) != 0 || len(r.sc.kept[4]) != 1 || r.sc.kept[4][0] != seeds[0] || len(r.keptSeeds) != 1 {
		t.Fatal("kept")
	}
	if _, ok := r.lam[Span{0, 4}]; ok {
		t.Fatal("lambda before the cut")
	}
	if _, ok := r.lam[Span{4, 8}]; !ok {
		t.Fatal("lambda after the cut")
	}
}
