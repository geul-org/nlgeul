//ff:func feature=parse type=reader control=sequence
//ff:what TestStreamNextWaits — a rule that reads the word after a candidate (waitingComplementRule) makes Next read on to the predicate before it judges: the tied 주격 and 보격 이 of 학생이 change places before 되었다 (보격 first, its preference up by exactly 2, 주격 and the new word as they were, no candidate taken away) and keep them before 밥을 먹었다, which is read with it up to 먹었다 and no further; a sentence that ends with no predicate is judged all the same (학생이 alone keeps 주격 first, the rule asked and fallen)
package parse

import "testing"

func TestStreamNextWaits(t *testing.T) {
	const subj, comp = "학생[열쇠:체언]+이[주격:이]", "학생[열쇠:체언]+이[보격:이]"
	plain, _ := sentenceStream(t, "학생이 되었다", false).Next()
	was := plain.Nodes[0]
	st := sentenceStream(t, "학생이 되었다", false, waitingComplementRule())
	seg, ok := st.Next()
	if !ok || len(st.queue) != 1 || !st.queue[0].judged || segmentLine(seg) != "학생이[0-3]" {
		t.Fatal("read on to 되었다 before 학생이 is sent", len(st.queue))
	}
	now := seg.Nodes[0]
	if len(now.Contest.Branches) != len(was.Contest.Branches) || now.Contest.Branches[0].String() != comp || now.Contest.Branches[1].String() != subj {
		t.Fatal(now.Contest.Branches)
	}
	oldComp, _ := branchNamed(was, comp)
	oldSubj, _ := branchNamed(was, subj)
	if !almost(now.Contest.Branches[0].Pref, oldComp.Pref+2) || now.Contest.Branches[1].Pref != oldSubj.Pref || now.Contest.Branches[2].Pref != was.Contest.Branches[2].Pref {
		t.Fatal("only 보격 is raised, by the bonus")
	}

	far := sentenceStream(t, "학생이 밥을 먹었다 학생이", false, waitingComplementRule())
	seg, _ = far.Next()
	if len(far.queue) != 2 || begun(far) != 3 || seg.Nodes[0].Contest.Branches[0].String() != subj || seg.Nodes[0].Contest.Branches[1].Pref != oldComp.Pref {
		t.Fatal("before 밥을 the rule falls; read up to 먹었다", len(far.queue), begun(far))
	}

	alone := sentenceStream(t, "학생이", false, waitingComplementRule())
	seg, ok = alone.Next()
	if !ok || seg.Nodes[0].Contest.Branches[0].String() != subj || !alone.sent[0].judged {
		t.Fatal("judged at the end of the sentence")
	}
}
