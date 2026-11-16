//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPredAfterObjectRuleStream — read as a stream with words-object.tsv (the noun 보니 read before the verb 보 with the ending 니, by less than bonusOverCount): with no rule 보니 is the noun; with the rule 보니 right after an object (살림을) is 보 with the ending 니 and the trace says 통과, and 보니 after 살림도 stays the noun and the trace says 기각: whyNoObjectBefore
package token

import "testing"

func TestPredAfterObjectRuleStream(t *testing.T) {
	r := predAfterObjectRule()
	line := func(v, why string) string {
		return "  · 보[열쇠:동사]+니[맺음:니]\t" + rulePredAfterObject + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"살림을 보니 좋다", "보[열쇠:동사]+니[맺음:니]", line("1.00", "통과")},
		{"살림도 보니 좋다", "보니[열쇠:체언]", line("-1.00", "기각: "+whyNoObjectBefore)},
	} {
		if got, _ := ruledFirst(t, "words-object.tsv", c.sentence, "보니"); got != "보니[열쇠:체언]" {
			t.Errorf("%s: 보니 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-object.tsv", c.sentence, "보니", r)
		if got != c.ruled {
			t.Errorf("%s: 보니 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
