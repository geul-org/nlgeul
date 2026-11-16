//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestOutsideNounRuleStream — read as a stream with words-outside.tsv: with no rule 밖에 right after a closing quotation mark is the particle 밖에; with the rule it is the noun 밖 with 에 where no negative follows (“집”밖에 나갔다) and the trace says 통과, and stays the particle before 없다 (“집”밖에 없다) and the trace says 기각: whyNegNext
package token

import "testing"

func TestOutsideNounRuleStream(t *testing.T) {
	r := outsideNounRule()
	noun := "밖[열쇠:체언]+에[부사격:에]"
	line := func(v, why string) string {
		return "  · " + noun + "\t" + ruleOutsideNoun + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"“집”밖에 나갔다", noun, line("1.00", "통과")},
		{"“집”밖에 없다", "+밖에[보조사:밖에]", line("-1.00", "기각: "+whyNegNext)},
	} {
		if got, _ := outsideFirst(t, c.sentence, "밖에"); got != "+밖에[보조사:밖에]" {
			t.Errorf("%s: 밖에 with no rule is %s", c.sentence, got)
		}
		got, notes := outsideFirst(t, c.sentence, "밖에", r)
		if got != c.ruled {
			t.Errorf("%s: 밖에 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
