//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestApartParticleRuleVocative — read as a stream with words-apart.tsv, where the interjection 아 stands beside the vocative 아 written apart: at the head of a sentence the rule does not hold for the vocative (기각: whyVocativeHead) and 아 stays the interjection, and after a noun (나무 아) it holds (통과; the vocative particle gains bonusNewTaken), though with the values fitted on 2026-11-19 the interjection of the small table stays first
package token

import "testing"

func TestApartParticleRuleVocative(t *testing.T) {
	r := apartParticleRule()
	alone, particle := "아[열쇠:홀로]", "+아[호격:아]"
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"아 나무 자란다", alone, "  · " + particle + "\t" + ruleApartParticle + "\t-1.00\t기각: " + whyNotAfterNominal + ", " + whyHeadOrSign + ", " + whyVocativeHead},
		{"나무 아 자란다", alone, "  · " + particle + "\t" + ruleApartParticle + "\t1.00\t통과"},
	} {
		got, notes := apartFirst(t, c.sentence, "아", r)
		if got != c.ruled {
			t.Errorf("%s: 아 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
	if got, _ := apartFirst(t, "아 나무 자란다", "아", r); got != alone {
		t.Errorf("아 at the head: %s", got)
	}
}
