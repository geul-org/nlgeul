//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestTableDeterminerRuleStream — read as a stream with words-alone.tsv (the nominal 총 counted more than the determiner): with no rule the nominal comes first; with the rule 총 before a noun is the determiner and the trace says 통과, and 총 before a predicate stays the nominal and the trace says 기각: whyNoModifiableNext
package parse

import "testing"

func TestTableDeterminerRuleStream(t *testing.T) {
	ix := streamReader(t, "words-alone.tsv").ix
	r := tableDeterminerRule(ix.lex, ix.closed)
	line := func(v, why string) string {
		return "  · 총[열쇠:관형사]\t" + ruleTableDeterminer + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"총 비용이 좋다", "총[열쇠:관형사]", line("1.00", "통과")},
		{"비용이 총 좋다", "총[열쇠:체언]", line("-1.00", "기각: "+whyNoModifiableNext)},
	} {
		if got, _ := aloneFirst(t, c.sentence, "총"); got != "총[열쇠:체언]" {
			t.Errorf("%s: 총 with no rule is %s", c.sentence, got)
		}
		got, notes := aloneFirst(t, c.sentence, "총", r)
		if got != c.ruled {
			t.Errorf("%s: 총 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
