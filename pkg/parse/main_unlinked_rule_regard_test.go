//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestMainUnlinkedRuleRegard — the main verb 보 of 보건대 judged with the word before it given: after 관하여 (관하+여) the linking form does not count (its exception is defeated) and the rule stands with bonusMainFirst, after 먹어 it falls with the reason 앞 어절이 연결 꼴(어/아, 고)로 끝난다
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMainUnlinkedRuleRegard(t *testing.T) {
	main := "보[열쇠:동사]+건대[맺음:건대]"
	words, _ := linkBeforeWords(t, "보건대", false)
	cand := auxLinkedPick(t, words[0], main)
	st := sentenceStream(t, "", true, mainUnlinkedRule([]string{"보"}))
	for _, c := range []struct {
		before rule.Near
		add    float64
		note   string
	}{
		{rule.Near{Stem: "관하", Kind: "열쇠:동사", Last: "맺음:여"}, bonusMainFirst, "1.00\t통과"},
		{rule.Near{Stem: "먹", Kind: "열쇠:동사", Last: "맺음:어"}, 0, "-1.00\t기각: 앞 어절이 연결 꼴(어/아, 고)로 끝난다"},
	} {
		add, _, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{c.before}})
		if add != c.add || len(notes) != 1 || notes[0] != "  · "+main+"\t"+ruleMainUnlinked+"\t"+c.note {
			t.Fatalf("%+v: %v %q", c.before, add, notes)
		}
	}
}
