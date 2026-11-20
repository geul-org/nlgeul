//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestApartParticleRuleSigns — read as a stream with words-apart.tsv: 는 written onto a run of signs that stands for a name kept back (○○는, □□□는) is 늘다 with no rule and the particle with the rule, which the trace says passes; written onto a noun it is read in the noun's own word, and 는 written onto a word of letters stays out of the rule (나무는 has no lone 는)
package parse

import "testing"

func TestApartParticleRuleSigns(t *testing.T) {
	r := apartParticleRule()
	verb, particle := "늘[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]", "+는[보조사:는]"
	for _, sentence := range []string{"○○는 자란다", "□□□는 자란다"} {
		if got, _ := apartFirst(t, sentence, "는"); got != verb {
			t.Errorf("%s: 는 with no rule is %s", sentence, got)
		}
		got, notes := apartFirst(t, sentence, "는", r)
		if got != particle {
			t.Errorf("%s: 는 with the rule is %s, want %s", sentence, got, particle)
		}
		auxLinkedHas(t, notes, "  · "+particle+"\t"+ruleApartParticle+"\t1.00\t통과")
	}
	if got, _ := apartFirst(t, "나무는 자란다", "는", r); got != "" {
		t.Errorf("나무는: a lone 는 %s", got)
	}
}
