//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDativeContractRuleStream — read as a stream with words-dative.tsv (the verb 내 counted more than the noun): with no rule 내게 is the verb 내 with the ending 게; with the rule 내게 before 주다 is the pronoun with the particle 게 and the trace says 통과, and 내게 before 하다 stays the verb and the trace says 기각: whyCauseNext
package token

import "testing"

func TestDativeContractRuleStream(t *testing.T) {
	r := dativeContractRule()
	line := func(v, why string) string {
		return "  · 내[열쇠:체언]+게[부사격:에게]\t" + ruleDativeContract + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"내게 책을 주었다", "내[열쇠:체언]+게[부사격:에게]", line("1.00", "통과")},
		{"책을 내게 하였다", "내[열쇠:동사]+게[맺음:게]", line("-1.00", "기각: "+whyCauseNext)},
	} {
		if got, _ := ruledFirst(t, "words-dative.tsv", c.sentence, "내게"); got != "내[열쇠:동사]+게[맺음:게]" {
			t.Errorf("%s: 내게 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-dative.tsv", c.sentence, "내게", r)
		if got != c.ruled {
			t.Errorf("%s: 내게 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
