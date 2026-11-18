//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAuxLinkedRuleGo — an auxiliary verb judged after -고: the auxiliary 주 of 주는 falls with the reason 앞 어절이 -고로 끝난다 and adds nothing, the auxiliaries 있 of 있는 and 보 of 보니 stand (the exception is defeated: the stem is one -고 is given before) and add bonusOverCount, and 주 after the 어 form (먹어) stands
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxLinkedRuleGo(t *testing.T) {
	st := sentenceStream(t, "", true, auxLinkedRule([]string{"보", "있", "주"}))
	for _, c := range []struct {
		in, cand, last string
		add            float64
		note           string
	}{
		{"주는", "주[보조 용언:주]+는[꾸밈:는]", "맺음:고", 0, "-1.00\t기각: " + whyGoNotAux},
		{"있는", "있[보조 용언:있]+는[꾸밈:는]", "맺음:고", bonusOverCount, "1.00\t통과"},
		{"보니", "보[보조 용언:보]+니[맺음:니]", "맺음:고", bonusOverCount, "1.00\t통과"},
		{"주는", "주[보조 용언:주]+는[꾸밈:는]", "맺음:어", bonusOverCount, "1.00\t통과"},
	} {
		words, _ := linkBeforeWords(t, c.in, false)
		cand := auxLinkedPick(t, words[0], c.cand)
		add, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{{Stem: "받", Kind: "열쇠:동사", Last: c.last, Marks: c.last}}})
		if add != c.add || len(notes) != 1 || notes[0] != "  · "+c.cand+"\t"+ruleAuxLinked+"\t"+c.note {
			t.Fatalf("%q after %s: %v %q", c.in, c.last, add, notes)
		}
	}
}
