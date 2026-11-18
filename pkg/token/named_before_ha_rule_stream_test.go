//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNamedBeforeHaRuleStream — read as a stream with words-named-before-ha.tsv (the noun 서기 with a particle read before the verb 서 with -기 and the particle by less than bonusOverCount): with no rule 서기도 and 서기로 are the noun with the particle; with the rule, right before 한다, they are 서 with -기 and the particle and the trace says 통과, and 서기도 before 간다 stays the noun and the trace says 기각: whyNoDoNext
package token

import "testing"

func TestNamedBeforeHaRuleStream(t *testing.T) {
	r := namedBeforeHaRule()
	line := func(cand, v, why string) string {
		return "  · " + cand + "\t" + ruleNamedBeforeHa + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, word, bare, ruled, note string
	}{
		{"그가 서기도 한다", "서기도", "서기[열쇠:체언]+도[보조사:도]", "서[열쇠:동사]+기[이름:기]+도[보조사:도]", line("서[열쇠:동사]+기[이름:기]+도[보조사:도]", "1.00", "통과")},
		{"그가 서기로 한다", "서기로", "서기[열쇠:체언]+로[부사격:로]", "서[열쇠:동사]+기[이름:기]+로[부사격:로]", line("서[열쇠:동사]+기[이름:기]+로[부사격:로]", "1.00", "통과")},
		{"그가 서기도 간다", "서기도", "서기[열쇠:체언]+도[보조사:도]", "서기[열쇠:체언]+도[보조사:도]", line("서[열쇠:동사]+기[이름:기]+도[보조사:도]", "-1.00", "기각: "+whyNoDoNext)},
	} {
		if got, _ := ruledFirst(t, "words-named-before-ha.tsv", c.sentence, c.word); got != c.bare {
			t.Errorf("%s: %s with no rule is %s, want %s", c.sentence, c.word, got, c.bare)
		}
		got, notes := ruledFirst(t, "words-named-before-ha.tsv", c.sentence, c.word, r)
		if got != c.ruled {
			t.Errorf("%s: %s with the rule is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
