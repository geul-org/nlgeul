//ff:func feature=token type=helper control=sequence
//ff:what TestBranchesOf — the candidates of the states at one span, one for each way a word is printed, in the order made, each one word with its verdict, preference and number of birth; of two that print the same word the better by the order candidates have stands in the place of the first, the first when equal
package token

import "testing"

func TestBranchesOf(t *testing.T) {
	sc := scanOf("가나")
	root := sc.registry().root
	worse := addState(sc, []*path{root}, wordStep(2, "same", 1, -5))
	other := addState(sc, []*path{root}, wordStep(2, "other", 1, -4))
	better := addState(sc, []*path{root}, wordStep(2, "same", 1, -3))
	later := addState(sc, []*path{root}, wordStep(2, "same", 1, -4))
	got := branchesOf(sc.ix, []*path{worse, other, better, later})
	if len(got) != 2 || got[0].String() != "same" || got[0].Pref != -3 || got[1].String() != "other" || len(got[0].scores) != 1 {
		t.Fatal(got)
	}
	if got = branchesOf(sc.ix, nil); got != nil {
		t.Fatal("none")
	}
}
