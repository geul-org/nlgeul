//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPossessiveRuleStream — read as a stream with words-possessive.tsv (the interjection 네 counted more than the pronoun) and the rule of a word of a nominal and an interjection standing alone beside it: with no rule 네 is the interjection; with the rule 네 before a noun is the pronoun and the trace says 통과, 네 before a predicate stays the interjection and the trace says 기각: whyNoModifiableNext, and 네 before the unit 개 stays the interjection and the trace says 기각: whyUnitNext
package token

import "testing"

func TestPossessiveRuleStream(t *testing.T) {
	ix := streamReader(t, "words-possessive.tsv").ix
	rules := []sentenceRule{interjectionAloneRule(ix.lex, ix.closed), possessiveRule()}
	line := func(v, why string) string {
		return "  · 네[열쇠:체언]\t" + rulePossessive + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"네 집으로 가", "네[열쇠:체언]", line("1.00", "통과")},
		{"네 좋다", "네[열쇠:홀로]", line("-1.00", "기각: "+whyNoModifiableNext)},
		{"네 개 좋다", "네[열쇠:홀로]", line("-1.00", "기각: "+whyUnitNext)},
	} {
		if got, _ := ruledFirst(t, "words-possessive.tsv", c.sentence, "네", rules[0]); got != "네[열쇠:홀로]" {
			t.Errorf("%s: 네 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-possessive.tsv", c.sentence, "네", rules...)
		if got != c.ruled {
			t.Errorf("%s: 네 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
