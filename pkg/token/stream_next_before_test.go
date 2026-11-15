//ff:func feature=token type=reader control=sequence
//ff:what TestStreamNextBefore — a rule that reads only the words before a candidate (waits nil) does not make Next read on: 학생이 is sent before 되었다 is read, and 되었다 reads the word already sent (behind) as its first candidate stood when sent: 주격 이 (the rule falls, 되었다 as it was), or 보격 이 when another rule that reads nothing around raised it (the rule stands and 되었다 gains 3)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamNextBefore(t *testing.T) {
	plain := streamFirsts(sentenceStream(t, "학생이 되었다", false))
	pred := plain[1].Nodes[0].Contest.Branches[0]

	st := sentenceStream(t, "학생이 되었다", false, afterComplementRule())
	first, _ := st.Next()
	if begun(st) != 1 || len(st.queue) != 0 || first.Nodes[0].Contest.Branches[0].String() != "학생[열쇠:체언]+이[주격:이]" {
		t.Fatal("not read on", begun(st))
	}
	second, _ := st.Next()
	if got := second.Nodes[0].Contest.Branches[0]; got.Pref != pred.Pref || len(st.behind) != 2 || st.behind[1].Last != "주격:이" {
		t.Fatal("after 주격 the rule falls", got.Pref, pred.Pref)
	}

	raise := sentenceRuleOf("보격", 1, waitingComplementRule().about, nil, listed("보격", rule.ListSelfLast, []string{"보격:이"}), nil, "")
	both := sentenceStream(t, "학생이 되었다", false, raise, afterComplementRule())
	first, _ = both.Next()
	if begun(both) != 1 || first.Nodes[0].Contest.Branches[0].String() != "학생[열쇠:체언]+이[보격:이]" {
		t.Fatal("raised and sent alone")
	}
	second, _ = both.Next()
	if got := second.Nodes[0].Contest.Branches[0]; got.String() != pred.String() || !almost(got.Pref, pred.Pref+3) || !almost(second.Lambda, plain[1].Lambda+3) {
		t.Fatal("after 보격 the rule stands", got.Pref, pred.Pref)
	}
}
