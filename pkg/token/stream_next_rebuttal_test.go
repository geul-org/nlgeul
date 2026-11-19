//ff:func feature=token type=reader control=sequence
//ff:what TestStreamNextRebuttal — Next with a rule of rebuttal: 주격 이 of 학생이 before 되었다 stands behind 보격 이 with a third of its verdict and its preference as it was, no candidate is taken away, the trace line comes with the stretch, and Lambda (the tied first lambda) is as it was; the word after reads 보격 이 before it (되었다 gains 3 by afterComplementRule); NextHead cuts the candidates in the order after the rebuttal; two rebuttals multiply (a ninth)
package token

import (
	"slices"
	"testing"
)

func TestStreamNextRebuttal(t *testing.T) {
	const s = "학생이 되었다"
	const subj, comp = "학생[열쇠:체언]+이[주격:이]", "학생[열쇠:체언]+이[보격:이]"
	plain, _ := sentenceStream(t, s, false).Next()
	st := sentenceStream(t, s, true, subjectRebuttalRule(), afterComplementRule())
	seg, _ := st.Next()
	was, _ := branchNamed(plain.Nodes[0], subj)
	bs := seg.Nodes[0].Contest.Branches
	now, k := branchNamed(seg.Nodes[0], subj)
	if bs[0].String() != comp || k != 1 || len(bs) != len(plain.Nodes[0].Contest.Branches) || !almost(now.Verdict, was.Verdict/3) || now.Pref != was.Pref || !almost(seg.Lambda, plain.Lambda) {
		t.Fatal(bs, seg.Lambda, plain.Lambda)
	}
	if !slices.Contains(st.Notes(), "  · "+subj+"\t되 앞 주격\t1.00\t통과") {
		t.Fatal("the trace line of the rebuttal comes with the stretch")
	}
	next, _ := st.Next()
	if bare, _ := sentenceStream(t, "되었다", false).Next(); !almost(next.Nodes[0].Contest.Branches[0].Pref, bare.Nodes[0].Contest.Branches[0].Pref+3) {
		t.Fatal("the word after reads 보격 이 before it", next.Nodes[0].Contest.Branches[0].Pref)
	}
	two, _ := sentenceStream(t, s, false, subjectRebuttalRule()).NextHead(2)
	if hs := two.Nodes[0].Contest.Branches; len(hs) != 2 || hs[0].String() != comp || hs[1].String() != subj || hs[1].Verdict != now.Verdict {
		t.Fatal(hs)
	}
	twice := subjectRebuttalRule()
	twice.name = "둘째"
	down, _ := sentenceStream(t, s, false, subjectRebuttalRule(), twice).Next()
	if br, _ := branchNamed(down.Nodes[0], subj); !almost(br.Verdict, was.Verdict/9) || br.Pref != was.Pref {
		t.Fatal(br.Verdict, was.Verdict)
	}
}
