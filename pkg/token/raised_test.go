//ff:func feature=token type=helper control=sequence
//ff:what TestRaised — a candidate gains the sum of the lifts that name it as it is printed, both in its preference and in the score kept of its word, and the scores are a copy (the branch raised from is not touched); a lift that names another candidate adds nothing, and with nothing added the branch comes back as it was; a branch with no scores kept gains only in its preference
package token

import "testing"

func TestRaised(t *testing.T) {
	n := firstNode(t, "학생이")
	comp, _ := branchNamed(n, "학생[열쇠:체언]+이[보격:이]")
	if len(comp.scores) == 0 {
		t.Fatal("a candidate of a stream keeps its scores")
	}
	pref, score := comp.Pref, comp.scores[0].pref
	own := []lift{{span: n.Span, word: comp.String(), add: 2}, {span: n.Span, word: "학생[열쇠:체언]+이[주격:이]", add: 5}, {span: n.Span, word: comp.String(), add: 0.5}}
	up := raised(comp, own)
	if up.Pref != pref+2.5 || up.scores[0].pref != score+2.5 || !almost(up.Lambda(), comp.Lambda()+2.5) {
		t.Fatal(up.Pref, pref, up.scores[0].pref, score)
	}
	if comp.Pref != pref || comp.scores[0].pref != score {
		t.Fatal("the branch raised from is touched", comp.scores[0].pref)
	}
	if same := raised(comp, own[1:2]); same.Pref != pref || &same.scores[0] != &comp.scores[0] {
		t.Fatal("a lift of another candidate adds nothing")
	}
	bare := raised(Branch{Words: comp.Words, Pref: -1}, own)
	if bare.Pref != 1.5 || bare.scores != nil {
		t.Fatal(bare.Pref)
	}
}
