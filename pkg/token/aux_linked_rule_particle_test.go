//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAuxLinkedRuleParticle — an auxiliary verb judged after a word whose markers hold a linking closing and end in an auxiliary particle: the auxiliary 있 of 있는 after -고는 and 보 of 보니 after -어도 stand and add bonusOverCount; 있 after a noun with 는 (no linking closing) and after -고 with a case particle at its end fall with the reason 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다 among its reasons; 주 of 주는 after -고는 falls with the reason 앞 어절이 -고로 끝난다 among them (the trace names every exception that applied)
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxLinkedRuleParticle(t *testing.T) {
	st := sentenceStream(t, "", true, auxLinkedRule([]string{"보", "있", "주"}))
	for _, c := range []struct {
		in, cand string
		before   rule.Near
		add      float64
		verdict  string
		reason   string
	}{
		{"있는", "있[보조 용언:있]+는[꾸밈:는]", rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "보조사:는", Marks: "맺음:고+보조사:는"}, bonusOverCount, "1.00\t통과", ""},
		{"보니", "보[보조 용언:보]+니[맺음:니]", rule.Near{Stem: "읽", Kind: "열쇠:동사", Last: "보조사:도", Marks: "맺음:어+보조사:도"}, bonusOverCount, "1.00\t통과", ""},
		{"있는", "있[보조 용언:있]+는[꾸밈:는]", rule.Near{Stem: "집", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}, 0, "-1.00\t기각: ", whyNotLinked},
		{"있는", "있[보조 용언:있]+는[꾸밈:는]", rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "목적격:를", Marks: "맺음:고+목적격:를"}, 0, "-1.00\t기각: ", whyNotLinked},
		{"주는", "주[보조 용언:주]+는[꾸밈:는]", rule.Near{Stem: "받", Kind: "열쇠:동사", Last: "보조사:는", Marks: "맺음:고+보조사:는"}, 0, "-1.00\t기각: ", whyGoNotAux},
	} {
		words, _ := linkBeforeWords(t, c.in, false)
		cand := auxLinkedPick(t, words[0], c.cand)
		add, _, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{c.before}})
		if add != c.add || len(notes) != 1 || !strings.HasPrefix(notes[0], "  · "+c.cand+"\t"+ruleAuxLinked+"\t"+c.verdict) || !strings.Contains(notes[0], c.reason) {
			t.Fatalf("%q after %+v: %v %q", c.in, c.before, add, notes)
		}
	}
}
