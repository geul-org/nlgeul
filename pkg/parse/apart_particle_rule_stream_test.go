//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestApartParticleRuleStream — read as a stream with words-apart.tsv: with no rule 는 written apart is 늘다 with the modifying ㄴ; with the rule it is the particle 는 after a noun (나무 는), after a closing quotation mark (“나무” 는) and at the head of a sentence, and the trace says 통과, and after an adverb (빨리 는) it stays 늘다 and the trace says 기각: whyNotAfterNominal
package parse

import "testing"

func TestApartParticleRuleStream(t *testing.T) {
	r := apartParticleRule()
	verb, particle := "늘[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]", "+는[보조사:는]"
	line := func(v, why string) string {
		return "  · " + particle + "\t" + ruleApartParticle + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"나무 는 자란다", particle, line("1.00", "통과")},
		{"“나무” 는 자란다", particle, line("1.00", "통과")},
		{"는 자란다", particle, line("1.00", "통과")},
		{"빨리 는 자란다", verb, line("-1.00", "기각: "+whyNotAfterNominal)},
	} {
		if got, _ := apartFirst(t, c.sentence, "는"); got != verb {
			t.Errorf("%s: 는 with no rule is %s", c.sentence, got)
		}
		got, notes := apartFirst(t, c.sentence, "는", r)
		if got != c.ruled {
			t.Errorf("%s: 는 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
