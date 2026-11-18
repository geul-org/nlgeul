//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestHaOneRuleStream — read as a stream with words-ha-one.tsv (the determiner 한 before 하다 with ㄴ, as the word table counted again puts them): with no rule the determiner comes first; with the rule 한 after the object particle (일을 한), after -게 (오게 한) and written onto a noun (창백한) is 하다 with ㄴ and the trace says 통과, and 한 at the head of a sentence, right after an opening quotation mark and after the object particle right before a unit (일을 한 번) stays the determiner and the trace says 기각
package token

import "testing"

func TestHaOneRuleStream(t *testing.T) {
	r := haOneRule()
	det, ha := "한[관형사:한]", "하[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]"
	line := func(v, why string) string {
		return "  · " + ha + "\t" + ruleHaOne + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"일을 한 사람이 왔다", ha, line("1.00", "통과")},
		{"오게 한 사람이다", ha, line("1.00", "통과")},
		{"창백한 사람이다", ha, line("1.00", "통과")},
		{"한 사람이 왔다", det, line("-1.00", "기각: "+whyNoHaTakes)},
		{"“한 사람이다", det, line("-1.00", "기각: "+whyNoHaTakes+", "+whyGluedWord+", "+whyOpenBefore)},
		{"일을 한 번 했다", det, line("-1.00", "기각: "+whyUnitNext)},
	} {
		if got, _ := haOneFirst(t, c.sentence, "한"); got != det {
			t.Errorf("%s: 한 with no rule is %s, want %s", c.sentence, got, det)
		}
		got, notes := haOneFirst(t, c.sentence, "한", r)
		if got != c.ruled {
			t.Errorf("%s: 한 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
