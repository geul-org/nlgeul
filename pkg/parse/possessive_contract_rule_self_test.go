//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveContractRuleSelf — read as a stream with words-possessive.tsv: with no rule 제 is the nominal 제; with the rule 제 before a noun is 저 with 의 fused and the trace says 통과, 제 before a number stays the nominal (기각: whyNumberNext), and so does 제 before a predicate (기각: whyNoModifiableNext)
package parse

import "testing"

func TestPossessiveContractRuleSelf(t *testing.T) {
	const fused, plain = "저[열쇠:체언]+(의)[관형격:의]", "제[열쇠:체언]"
	r := possessiveContractRule()
	line := func(v, why string) string {
		return "  · " + fused + "\t" + rulePossessiveContract + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"제 집으로 가", fused, line("1.00", "통과")},
		{"제 3 집으로 가", plain, line("-1.00", "기각: "+whyNumberNext)},
		{"제 좋다", plain, line("-1.00", "기각: "+whyNoModifiableNext)},
	} {
		if got, _ := ruledFirst(t, "words-possessive.tsv", c.sentence, "제"); got != plain {
			t.Errorf("%s: 제 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-possessive.tsv", c.sentence, "제", r)
		if got != c.ruled {
			t.Errorf("%s: 제 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
