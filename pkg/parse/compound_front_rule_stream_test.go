//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCompoundFrontRuleStream — read as a stream with words-compound.tsv (the adverb 우선 counted more than the nominal): with no rule 우선 is the adverb; with the rule 우선 with the noun 출자 written onto it is the nominal and the trace says 통과, and 우선 written apart from 출자 stays the adverb and the trace says 기각: whyApartNext
package parse

import "testing"

func TestCompoundFrontRuleStream(t *testing.T) {
	r := compoundFrontRule(streamReader(t, "words-compound.tsv").ix.lex)
	line := func(v, why string) string {
		return "  · 우선[열쇠:체언]\t" + ruleCompoundFront + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"학교가 우선출자를 하다", "우선[열쇠:체언]", line("1.00", "통과")},
		{"학교가 우선 출자를 하다", "우선[열쇠:부사]", line("-1.00", "기각: "+whyApartNext)},
	} {
		if got, _ := ruledFirst(t, "words-compound.tsv", c.sentence, "우선"); got != "우선[열쇠:부사]" {
			t.Errorf("%s: 우선 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-compound.tsv", c.sentence, "우선", r)
		if got != c.ruled {
			t.Errorf("%s: 우선 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
