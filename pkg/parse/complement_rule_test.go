//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestComplementRule — a 이/가 candidate right before 되다 or 아니다 is read as the complement first: in 물이 얼음이 되었다 얼음이 has the 보격 first with its preference up by exactly bonusSameLetters over the unruled stream, and 물이 keeps the 주격 first with both preferences as they were, and the rule speaks of and waits for the 보격 candidate alone; in 그는 학생이 아니다 학생이 has the 보격 first; the trace line of the lifted candidate is 통과
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestComplementRule(t *testing.T) {
	complementLifted(t, "물이 얼음이 되었다", 1, "얼음[열쇠:체언]+이[보격:이]", "얼음[열쇠:체언]+이[주격:이]")
	complementLifted(t, "그는 학생이 아니다", 1, "학생[열쇠:체언]+이[보격:이]", "학생[열쇠:체언]+이[주격:이]")
	plain, _ := complementNodes(complementStream(t, "물이 얼음이 되었다", false, false))
	ruled, _ := complementNodes(complementStream(t, "물이 얼음이 되었다", false, true))
	if ruled[0].Contest.Branches[0].String() != "물[열쇠:체언]+이[주격:이]" {
		t.Fatal(ruled[0].Contest.Branches)
	}
	for i, br := range plain[0].Contest.Branches {
		now, at := branchNamed(ruled[0], br.String())
		if at < 0 || !almost(now.Pref, br.Pref) {
			t.Fatal("물이 is not lifted", i, ruled[0].Contest.Branches)
		}
	}
	r := complementRule()
	comp, subj := rule.Near{Stem: "학생", Last: "보격:이", Marks: "보격:이"}, rule.Near{Stem: "학생", Last: "주격:이", Marks: "주격:이"}
	if !r.about(comp) || !r.waits(comp) || r.about(subj) || r.waits(subj) || r.bonus != bonusSameLetters {
		t.Fatal("the rule speaks of, and waits for, the complement candidate alone")
	}
}
