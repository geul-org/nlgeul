//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestLinkBeforeNegRuleStream — read as a stream with words-neg.tsv (the noun 가지 counted more than the verb 가): with no rule 가지 is the noun; with the rule 가지 before 않다 is 가 with the ending 지 and the trace says 통과, and 가지 before 좋다 stays the noun and the trace says 기각: whyNoNegNext
package parse

import "testing"

func TestLinkBeforeNegRuleStream(t *testing.T) {
	r := linkBeforeNegRule()
	line := func(v, why string) string {
		return "  · 가[열쇠:동사]+지[맺음:지]\t" + ruleLinkBeforeNeg + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"가지 않는다", "가[열쇠:동사]+지[맺음:지]", line("1.00", "통과")},
		{"가지 좋다", "가지[열쇠:체언]", line("-1.00", "기각: "+whyNoNegNext)},
	} {
		if got, _ := ruledFirst(t, "words-neg.tsv", c.sentence, "가지"); got != "가지[열쇠:체언]" {
			t.Errorf("%s: 가지 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-neg.tsv", c.sentence, "가지", r)
		if got != c.ruled {
			t.Errorf("%s: 가지 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
