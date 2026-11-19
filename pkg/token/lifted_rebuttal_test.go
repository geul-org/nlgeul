//ff:func feature=token type=helper control=sequence
//ff:what TestLiftedRebuttal — a node with a rebuttal of its own span against its first candidate: the tied 주격 이 goes behind 보격 이 with its verdict lowered and its preference as it was, no candidate is taken away, and the node it was made from keeps its candidates, their verdicts and their order
package token

import "testing"

func TestLiftedRebuttal(t *testing.T) {
	n := firstNode(t, "학생이")
	const subj, comp = "학생[열쇠:체언]+이[주격:이]", "학생[열쇠:체언]+이[보격:이]"
	was, i := branchNamed(n, subj)
	if i != 0 {
		t.Fatal("the tied 주격 is first before the rebuttal", n.Contest.Branches)
	}
	down := lifted(n, []lift{{span: n.Span, word: subj, force: forceUngrammatical}})
	if len(down.Contest.Branches) != len(n.Contest.Branches) || down.Contest.Branches[0].String() != comp {
		t.Fatal(down.Contest.Branches)
	}
	now, k := branchNamed(down, subj)
	if k == 0 || !almost(now.Verdict, was.Verdict*(2/(1+forceUngrammatical) - 1)) || now.Pref != was.Pref {
		t.Fatal(k, now.Verdict, now.Pref)
	}
	if n.Contest.Branches[0].String() != subj || n.Contest.Branches[0].Verdict != was.Verdict {
		t.Fatal("the node lifted from is touched")
	}
}
