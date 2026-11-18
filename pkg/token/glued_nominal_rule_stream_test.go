//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestGluedNominalRuleStream — read as a stream with words-compound.tsv (the determiner 전 counted more than the nominal): with no rule 전 is the determiner; with the rule 전 written onto the noun 발병 is the nominal and the trace says 통과, and 전 written apart from 발병 stays the determiner and the trace says 기각: whyApartBefore
package token

import "testing"

func TestGluedNominalRuleStream(t *testing.T) {
	ix := streamReader(t, "words-compound.tsv").ix
	r := gluedNominalRule(ix.lex, ix.closed)
	line := func(v, why string) string {
		return "  · 전[열쇠:체언]\t" + ruleGluedNominal + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"발병전 학교에 가다", "전[열쇠:체언]", line("1.00", "통과")},
		{"발병 전 학교에 가다", "전[열쇠:관형사]", line("-1.00", "기각: "+whyApartBefore)},
	} {
		if got, _ := ruledFirst(t, "words-compound.tsv", c.sentence, "전"); got != "전[열쇠:관형사]" {
			t.Errorf("%s: 전 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-compound.tsv", c.sentence, "전", r)
		if got != c.ruled {
			t.Errorf("%s: 전 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
