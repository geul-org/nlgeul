//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveContractRuleYou — read as a stream with words-possessive.tsv (the interjection 네 counted more than the pronoun) and the rule of a word of a nominal and an interjection standing alone beside it: with no rule 네 is the interjection; with the rule 네 before a noun is 너 with 의 fused and the trace says 통과, 네 before a predicate stays the interjection (기각: whyNoModifiableNext), and so does 네 before a unit (기각: whyUnitNext)
package parse

import "testing"

func TestPossessiveContractRuleYou(t *testing.T) {
	const fused, alone = "너[열쇠:체언]+(의)[관형격:의]", "네[열쇠:홀로]"
	ix := streamReader(t, "words-possessive.tsv").ix
	rules := []sentenceRule{interjectionAloneRule(ix.lex, ix.closed), possessiveContractRule()}
	line := func(v, why string) string {
		return "  · " + fused + "\t" + rulePossessiveContract + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"네 집으로 가", fused, line("1.00", "통과")},
		{"네 좋다", alone, line("-1.00", "기각: "+whyNoModifiableNext)},
		{"네 개 좋다", alone, line("-1.00", "기각: "+whyUnitNext)},
	} {
		if got, _ := ruledFirst(t, "words-possessive.tsv", c.sentence, "네", rules[0]); got != alone {
			t.Errorf("%s: 네 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-possessive.tsv", c.sentence, "네", rules...)
		if got != c.ruled {
			t.Errorf("%s: 네 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
