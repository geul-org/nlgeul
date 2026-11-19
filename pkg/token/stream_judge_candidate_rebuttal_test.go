//ff:func feature=token type=rule control=sequence
//ff:what TestStreamJudgeCandidateRebuttal — a rule of rebuttal gives its force, not a bonus, when its claim stands (주격 이 before 되), with the trace line in the form of every rule (통과); where an exception defeats it (before 밥) it gives no force and the line 기각 with the reason; two rebuttals that stand give both forces, and a rule of preference that stands beside them still gives its bonus; with no trace the forces are the same and no line is written
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamJudgeCandidateRebuttal(t *testing.T) {
	n := firstNode(t, "학생이")
	subj, _ := branchNamed(n, "학생[열쇠:체언]+이[주격:이]")
	st := sentenceStream(t, "", true, subjectRebuttalRule())
	before := rule.Around{Self: nearOf(subj), After: []rule.Near{{Stem: "되", Pred: true}}}
	add, forces, notes := st.judgeCandidate(subj, before)
	if add != 0 || len(forces) != 1 || forces[0] != forceUngrammatical || len(notes) != 1 || notes[0] != "  · 학생[열쇠:체언]+이[주격:이]\t되 앞 주격\t1.00\t통과" {
		t.Fatalf("%v %v %q", add, forces, notes)
	}
	add, forces, notes = st.judgeCandidate(subj, rule.Around{Self: nearOf(subj), After: []rule.Near{{Stem: "밥"}}})
	if add != 0 || forces != nil || len(notes) != 1 || notes[0] != "  · 학생[열쇠:체언]+이[주격:이]\t되 앞 주격\t-1.00\t기각: 뒤 어절이 되가 아니다" {
		t.Fatalf("%v %v %q", add, forces, notes)
	}
	again := subjectRebuttalRule()
	again.name = "되 앞 주격 둘째"
	prefer := sentenceRuleOf("주격", 0.5, again.about, nil, listed("주격", rule.ListSelfLast, []string{"주격:이"}), nil, "")
	st.rules = append(st.rules, again, prefer)
	if add, forces, notes := st.judgeCandidate(subj, before); add != 0.5 || len(forces) != 2 || forces[1] != forceUngrammatical || len(notes) != 3 {
		t.Fatal(add, forces, notes)
	}
	quiet := sentenceStream(t, "", false, subjectRebuttalRule())
	if _, forces, notes := quiet.judgeCandidate(subj, before); len(forces) != 1 || notes != nil {
		t.Fatal(forces, notes)
	}
}
