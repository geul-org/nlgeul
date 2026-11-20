//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestMainUnlinkedRuleGo — a main predicate judged after -고: the main verb 주 of 주는 stands with bonusMainFirst (the linking form does not count before a stem -고 is not given before), while the main 있 of 있는 and the main 보 of 보니 fall with the reason 앞 어절이 연결 꼴(어/아, 고)로 끝난다 and add nothing, as 주 does after the 어 form (먹어)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMainUnlinkedRuleGo(t *testing.T) {
	st := sentenceStream(t, "", true, mainUnlinkedRule([]string{"보", "있", "주"}))
	for _, c := range []struct {
		in, cand, last string
		add            float64
		note           string
	}{
		{"주는", "주[열쇠:동사]+는[꾸밈:는]", "맺음:고", bonusMainFirst, "1.00\t통과"},
		{"있는", "있[열쇠:형용사]+는[꾸밈:는]", "맺음:고", 0, "-1.00\t기각: " + whyLinked},
		{"보니", "보[열쇠:동사]+니[맺음:니]", "맺음:고", 0, "-1.00\t기각: " + whyLinked},
		{"주는", "주[열쇠:동사]+는[꾸밈:는]", "맺음:어", 0, "-1.00\t기각: " + whyLinked},
	} {
		words, _ := linkBeforeWords(t, c.in, false)
		cand := auxLinkedPick(t, words[0], c.cand)
		add, _, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{{Stem: "받", Kind: "열쇠:동사", Last: c.last, Marks: c.last}}})
		if add != c.add || len(notes) != 1 || notes[0] != "  · "+c.cand+"\t"+ruleMainUnlinked+"\t"+c.note {
			t.Fatalf("%q after %s: %v %q", c.in, c.last, add, notes)
		}
	}
}
