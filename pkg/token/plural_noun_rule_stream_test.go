//ff:func feature=token type=rule control=sequence
//ff:what TestPluralNounRuleStream — read as a stream with words-compound.tsv (the adverb 모두 counted more than the nominal): with no rule 모두들 is the adverb with 들; with the rule it is the nominal with its plural suffix and the trace says 통과
package token

import "testing"

func TestPluralNounRuleStream(t *testing.T) {
	r := pluralNounRule(streamReader(t, "words-compound.tsv").ix.lex)
	if got, _ := ruledFirst(t, "words-compound.tsv", "모두들 가다", "모두들"); got != "모두[열쇠:부사]+들[접미:들]" {
		t.Errorf("모두들 with no rule is %s", got)
	}
	got, notes := ruledFirst(t, "words-compound.tsv", "모두들 가다", "모두들", r)
	if got != "모두[열쇠:체언]+들[접미:들]" {
		t.Errorf("모두들 with the rule is %s", got)
	}
	auxLinkedHas(t, notes, "  · 모두[열쇠:체언]+들[접미:들]\t"+rulePluralNoun+"\t1.00\t통과")
}
