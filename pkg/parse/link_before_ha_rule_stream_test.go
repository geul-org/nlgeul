//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestLinkBeforeHaRuleStream — read as a stream with words-link-before-ha.tsv (the noun 내야 read before the verb 내 with -어야 by less than bonusOverCount): with no rule 내야 is the noun; with the rule 내야 right before 한다 and 된다 is 내 with -어야 and the trace says 통과, and 내야 before 간다 stays the noun and the trace says 기각: whyNoHaNext
package parse

import "testing"

func TestLinkBeforeHaRuleStream(t *testing.T) {
	r := linkBeforeHaRule()
	verb := "내[열쇠:동사]+(어)야[맺음:어야(합침)]"
	line := func(v, why string) string {
		return "  · " + verb + "\t" + ruleLinkBeforeHa + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"그가 내야 한다", verb, line("1.00", "통과")},
		{"그가 내야 된다", verb, line("1.00", "통과")},
		{"그가 내야 간다", "내야[열쇠:체언]", line("-1.00", "기각: "+whyNoHaNext)},
	} {
		if got, _ := ruledFirst(t, "words-link-before-ha.tsv", c.sentence, "내야"); got != "내야[열쇠:체언]" {
			t.Errorf("%s: 내야 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-link-before-ha.tsv", c.sentence, "내야", r)
		if got != c.ruled {
			t.Errorf("%s: 내야 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
