//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestTableDeterminerRulePrior — read as a stream with words-prior.tsv (전 the determiner counted more than 전 the noun) and both tableDeterminerRule and priorNounRule: 전 after a -기 form, a number with its unit and a bare noun of an action is the noun and the determiner's rule says 기각 with the reason of priorExcept, and 전 after another noun stays the determiner with the determiner's rule 통과
package parse

import "testing"

func TestTableDeterminerRulePrior(t *testing.T) {
	ix := streamReader(t, "words-prior.tsv").ix
	rules := []sentenceRule{tableDeterminerRule(ix.lex, ix.closed), priorNounRule(ix.lex)}
	line := func(v, why string) string {
		return "  · 전[열쇠:관형사]\t" + ruleTableDeterminer + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"먹기 전 사람", "전[열쇠:체언]", line("-1.00", "기각: "+whyPriorBefore)},
		{"3년 전 사람", "전[열쇠:체언]", line("-1.00", "기각: "+whyPriorBefore)},
		{"발병 전 사람", "전[열쇠:체언]", line("-1.00", "기각: "+whyActionBefore)},
		{"사람 전 사람", "전[열쇠:관형사]", line("1.00", "통과")},
	} {
		got, notes := ruledFirst(t, "words-prior.tsv", c.sentence, "전", rules...)
		if got != c.ruled {
			t.Errorf("%s: 전 with the rules is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
