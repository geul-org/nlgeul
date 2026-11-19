//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMainUnlinkedRuleWhile — a main predicate judged after a word that ends in 다 or 다가: the main verb 보 of 보니 (니, 니까, 면: whileClosings) falls with the reason 앞 어절이 다·다가로 끝난다 and adds nothing, while the main verb 보 of 보아 and the main 있 of 있으니 stand with bonusMainFirst (the exception is defeated: not 니·니까·면, not the stem 보)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMainUnlinkedRuleWhile(t *testing.T) {
	st := sentenceStream(t, "", true, mainUnlinkedRule([]string{"보", "있"}))
	for _, c := range []struct {
		in, cand, last string
		add            float64
		note           string
	}{
		{"보니", "보[열쇠:동사]+니[맺음:니]", "맺음:다", 0, "-1.00\t기각: 앞 어절이 다·다가로 끝난다"},
		{"보니", "보[열쇠:동사]+니[맺음:니]", "맺음:다가", 0, "-1.00\t기각: 앞 어절이 다·다가로 끝난다"},
		{"보아", "보[열쇠:동사]+아[맺음:아]", "맺음:다", bonusMainFirst, "1.00\t통과"},
		{"있으니", "있[열쇠:형용사]+으[이음:으]+니[맺음:니]", "맺음:다", bonusMainFirst, "1.00\t통과"},
	} {
		words, _ := linkBeforeWords(t, c.in, false)
		cand := auxLinkedPick(t, words[0], c.cand)
		add, _, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{{Stem: "먹", Last: c.last}}})
		if add != c.add || len(notes) != 1 || notes[0] != "  · "+c.cand+"\t"+ruleMainUnlinked+"\t"+c.note {
			t.Fatalf("%q after %s: %v %q", c.in, c.last, add, notes)
		}
	}
}
