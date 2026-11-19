//ff:func feature=token type=rule control=sequence
//ff:what TestStreamJudgeCandidate — a candidate is judged by the rules that speak of it: 보격 이 before 되 gains the bonus with the trace line 통과, before 밥 gains nothing with the line 기각 and the reason of the exception, and gains the bonuses of two rules that stand added; a candidate no rule speaks of gains nothing and has no line; with no trace the gain is the same and no line is written
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamJudgeCandidate(t *testing.T) {
	n := firstNode(t, "학생이")
	comp, _ := branchNamed(n, "학생[열쇠:체언]+이[보격:이]")
	subj, _ := branchNamed(n, "학생[열쇠:체언]+이[주격:이]")
	st := sentenceStream(t, "", true, waitingComplementRule())
	before := rule.Around{Self: nearOf(comp), After: []rule.Near{{Stem: "되", Pred: true}}}
	add, _, notes := st.judgeCandidate(comp, before)
	if add != 2 || len(notes) != 1 || notes[0] != "  · 학생[열쇠:체언]+이[보격:이]\t보격 뒤 되\t1.00\t통과" {
		t.Fatalf("%v %q", add, notes)
	}
	add, _, notes = st.judgeCandidate(comp, rule.Around{Self: nearOf(comp), After: []rule.Near{{Stem: "밥"}}})
	if add != 0 || len(notes) != 1 || notes[0] != "  · 학생[열쇠:체언]+이[보격:이]\t보격 뒤 되\t-1.00\t기각: 뒤 어절이 되가 아니다" {
		t.Fatalf("%v %q", add, notes)
	}
	if add, _, notes := st.judgeCandidate(subj, rule.Around{Self: nearOf(subj), After: before.After}); add != 0 || notes != nil {
		t.Fatal("no rule speaks of 주격", add, notes)
	}
	st.rules = append(st.rules, sentenceRuleOf("보격", 0.5, waitingComplementRule().about, nil, listed("보격", rule.ListSelfLast, []string{"보격:이"}), nil, ""))
	if add, _, notes := st.judgeCandidate(comp, before); add != 2.5 || len(notes) != 2 {
		t.Fatal(add, notes)
	}
	quiet := sentenceStream(t, "", false, waitingComplementRule())
	if add, _, notes := quiet.judgeCandidate(comp, before); add != 2 || notes != nil {
		t.Fatal(add, notes)
	}
}
