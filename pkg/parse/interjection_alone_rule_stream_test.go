//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestInterjectionAloneRuleStream — read as a stream with words-alone.tsv (the nominal 흠 counted more than the interjection): with no rule the nominal comes first; with the rule 흠 before a comma is the interjection and the trace says 통과, and 흠 after a modifying form stays the nominal and the trace says 기각: whyModifierBefore
package parse

import "testing"

func TestInterjectionAloneRuleStream(t *testing.T) {
	ix := streamReader(t, "words-alone.tsv").ix
	r := interjectionAloneRule(ix.lex, ix.closed)
	line := func(v, why string) string {
		return "  · 흠[열쇠:홀로]\t" + ruleInterjectionAlone + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"흠, 좋다", "흠[열쇠:홀로]", line("1.00", "통과")},
		{"좋은 흠", "흠[열쇠:체언]", line("-1.00", "기각: "+whyModifierBefore)},
	} {
		if got, _ := aloneFirst(t, c.sentence, "흠"); got != "흠[열쇠:체언]" {
			t.Errorf("%s: 흠 with no rule is %s", c.sentence, got)
		}
		got, notes := aloneFirst(t, c.sentence, "흠", r)
		if got != c.ruled {
			t.Errorf("%s: 흠 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
