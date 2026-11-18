//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPriorNounRuleStream — read as a stream with words-prior.tsv (전 the determiner read before 전 the noun, by less than bonusOverCount): with no rule 전 is the determiner; with the rule 전 after a -기 form (먹기) and after a number with its unit (3년) and after a noun of an action (발병) is the noun and the trace says 통과, and 전 after a noun stays the determiner and the trace says 기각: whyNoPrior
package token

import "testing"

func TestPriorNounRuleStream(t *testing.T) {
	r := priorNounRule(streamReader(t, "words-prior.tsv").ix.lex)
	noun := "전[열쇠:체언]"
	line := func(v, why string) string {
		return "  · " + noun + "\t" + rulePriorNoun + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"먹기 전 사람", noun, line("1.00", "통과")},
		{"3년 전 사람", noun, line("1.00", "통과")},
		{"발병 전 사람", noun, line("1.00", "통과")},
		{"사람 전 사람", "전[열쇠:관형사]", line("-1.00", "기각: "+whyNoPrior)},
	} {
		if got, _ := ruledFirst(t, "words-prior.tsv", c.sentence, "전"); got != "전[열쇠:관형사]" {
			t.Errorf("%s: 전 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-prior.tsv", c.sentence, "전", r)
		if got != c.ruled {
			t.Errorf("%s: 전 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
