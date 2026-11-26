//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveContractRuleStream — read as a stream with words-possessive-contract.tsv: with no rule 내 is the nominal 내; with the rule 내 before a noun is 나 with 의 fused and the trace says 통과, 내 right after a nominal with no particle stays the nominal (기각: whyAfterNominal), and 내 before a predicate is not carried (기각: whyNoModifiableNext)
package parse

import "testing"

func TestPossessiveContractRuleStream(t *testing.T) {
	const fused, plain = "나[열쇠:체언]+(의)[관형격:의]", "내[열쇠:체언]"
	r := possessiveContractRule()
	line := func(v, why string) string {
		return "  · " + fused + "\t" + rulePossessiveContract + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"내 말을 들었다", fused, line("1.00", "통과")},
		{"기간 내 제출을 하였다", plain, line("-1.00", "기각: "+whyAfterNominal)},
		{"돈을 내 버렸다", plain, line("-1.00", "기각: "+whyNoModifiableNext)},
	} {
		if got, _ := ruledFirst(t, "words-possessive-contract.tsv", c.sentence, "내"); got != plain {
			t.Errorf("%s: 내 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-possessive-contract.tsv", c.sentence, "내", r)
		if got != c.ruled {
			t.Errorf("%s: 내 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
