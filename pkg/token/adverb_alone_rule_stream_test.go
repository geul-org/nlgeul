//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdverbAloneRuleStream — read as a stream with words-alone.tsv (the nominals 모두·우선 counted more than the adverbs): with no rule the nominal comes first; with the rule 모두 at the head before a noun with a particle is the adverb and the trace says 통과, and 모두 after the adnominal 의, 모두 written onto a word and 우선 with a noun written onto it stay the nominal and the trace says 기각 with the reason
package token

import "testing"

func TestAdverbAloneRuleStream(t *testing.T) {
	r := adverbAloneRule(streamReader(t, "words-alone.tsv").ix.lex)
	line := func(stem, v, why string) string {
		return "  · " + stem + "[열쇠:부사]\t" + ruleAdverbAlone + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, word, ruled, note string
	}{
		{"모두 학교에 갔다", "모두", "모두[열쇠:부사]", line("모두", "1.00", "통과")},
		{"학교의 모두 비용", "모두", "모두[열쇠:체언]", line("모두", "-1.00", "기각: "+whyModifierBefore)},
		{"학교에모두 갔다", "모두", "모두[열쇠:체언]", line("모두", "-1.00", "기각: "+whyGluedWord)},
		{"우선출자를 했다", "우선", "우선[열쇠:체언]", line("우선", "-1.00", "기각: "+whyGluedNext)},
	} {
		if got, _ := aloneFirst(t, c.sentence, c.word); got != c.word+"[열쇠:체언]" {
			t.Errorf("%s: %s with no rule is %s", c.sentence, c.word, got)
		}
		got, notes := aloneFirst(t, c.sentence, c.word, r)
		if got != c.ruled {
			t.Errorf("%s: %s with the rule is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
