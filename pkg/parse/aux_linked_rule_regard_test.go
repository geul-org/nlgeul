//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAuxLinkedRuleRegard — the auxiliary 보 of 보건대 judged with the word before it given: after 관하여, 대하여 and 통해 (regardStems) the rule falls with the reason whyRegard and adds nothing, after another verb's 어/아 form (먹어) it stands and adds bonusOverCount; read with words-link-before-aux.tsv, 점에 관하여 보건대 puts the main verb 보 first and 먹어 보건대 the auxiliary
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxLinkedRuleRegard(t *testing.T) {
	aux, main := "보[보조 용언:보]+건대[맺음:건대]", "보[열쇠:동사]+건대[맺음:건대]"
	words, _ := linkBeforeWords(t, "보건대", false)
	cand := auxLinkedPick(t, words[0], aux)
	st := sentenceStream(t, "", true, auxLinkedRule([]string{"보"}))
	for _, c := range []struct {
		before rule.Near
		add    float64
		note   string
	}{
		{rule.Near{Stem: "관하", Kind: "열쇠:동사", Last: "맺음:여"}, 0, "-1.00\t기각: " + whyRegard},
		{rule.Near{Stem: "대하", Kind: "열쇠:동사", Last: "맺음:여"}, 0, "-1.00\t기각: " + whyRegard},
		{rule.Near{Stem: "통하", Kind: "열쇠:동사", Last: "맺음:어(합침)"}, 0, "-1.00\t기각: " + whyRegard},
		{rule.Near{Stem: "먹", Kind: "열쇠:동사", Last: "맺음:어"}, bonusOverCount, "1.00\t통과"},
	} {
		add, _, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{c.before}})
		if add != c.add || len(notes) != 1 || notes[0] != "  · "+aux+"\t"+ruleAuxLinked+"\t"+c.note {
			t.Fatalf("%+v: %v %q", c.before, add, notes)
		}
	}
	for in, first := range map[string]string{"점에 관하여 보건대": main, "먹어 보건대": aux} {
		got, _ := linkBeforeWords(t, in, true)
		if last := got[len(got)-1]; last[0].String() != first {
			t.Fatalf("%q: %q first", in, last[0].String())
		}
	}
}
