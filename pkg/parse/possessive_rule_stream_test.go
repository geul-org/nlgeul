//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveRuleStream — read as a stream with words-possessive.tsv: 니 standing alone is the nominal with and without the rule (the table has no other reading above it); with the rule the trace says 통과 before a noun, 기각 before a predicate (whyNoModifiableNext) and 기각 before a unit (whyUnitNext)
package parse

import "testing"

func TestPossessiveRuleStream(t *testing.T) {
	r := possessiveRule()
	line := func(v, why string) string {
		return "  · 니[열쇠:체언]\t" + rulePossessive + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, note string
	}{
		{"니 집으로 가", line("1.00", "통과")},
		{"니 좋다", line("-1.00", "기각: "+whyNoModifiableNext)},
		{"니 개 좋다", line("-1.00", "기각: "+whyUnitNext)},
	} {
		got, notes := ruledFirst(t, "words-possessive.tsv", c.sentence, "니", r)
		if got != "니[열쇠:체언]" {
			t.Errorf("%s: 니 with the rule is %s", c.sentence, got)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
