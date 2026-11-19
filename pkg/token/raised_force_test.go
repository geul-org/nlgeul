//ff:func feature=token type=helper control=sequence
//ff:what TestRaisedForce — a rebuttal that names a candidate multiplies its verdict by 2/(1+force)-1 in the branch and in the score kept of its word and leaves its preference as it is (forceUngrammatical: by 2/(1+forceUngrammatical)-1, Lambda times its log10 less lambda); two multiply (by its square), with what a lift adds added beside; the scores are a copy; a rebuttal of another candidate and a force of 0 change nothing
package token

import (
	"math"
	"testing"
)

func TestRaisedForce(t *testing.T) {
	n := firstNode(t, "학생이")
	subj, _ := branchNamed(n, "학생[열쇠:체언]+이[주격:이]")
	verdict, pref, score := subj.Verdict, subj.Pref, subj.scores[0].verdict
	m := 2/(1+forceUngrammatical) - 1
	one := lift{span: n.Span, word: subj.String(), force: forceUngrammatical}
	down := raised(subj, []lift{one})
	if !almost(down.Verdict, verdict*m) || !almost(down.scores[0].verdict, score*m) || down.Pref != pref || down.scores[0].pref != subj.scores[0].pref {
		t.Fatal(down.Verdict, down.Pref, down.scores[0])
	}
	if !almost(down.Lambda(), subj.Lambda()+Lambda*math.Log10(m)) {
		t.Fatal(down.Lambda(), subj.Lambda())
	}
	if subj.Verdict != verdict || subj.scores[0].verdict != score {
		t.Fatal("the branch lowered from is touched")
	}
	two := raised(subj, []lift{one, {span: n.Span, word: subj.String(), add: 1}, one})
	if !almost(two.Verdict, verdict*m*m) || !almost(two.scores[0].verdict, score*m*m) || two.Pref != pref+1 {
		t.Fatal(two.Verdict, two.Pref)
	}
	other := lift{span: n.Span, word: "학생[열쇠:체언]+이[보격:이]", force: forceUngrammatical}
	if same := raised(subj, []lift{other, {span: n.Span, word: subj.String()}}); same.Verdict != verdict || &same.scores[0] != &subj.scores[0] {
		t.Fatal("a rebuttal of another candidate or a force of 0 changes nothing")
	}
}
