//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestOpenMarkerRuleStream — read as a stream with words-open-marker.tsv: with no rule 을 right after an opening bracket or quotation mark is the object particle; with the rule it is read as the new word 을 (the particle's rebuttal 통과)
package parse

import "testing"

func TestOpenMarkerRuleStream(t *testing.T) {
	const particle = "+을[목적격:을]"
	for _, sentence := range []string{"약정(을)은 끝났다.", "\"을 끝났다."} {
		if got, _ := ruledFirst(t, "words-open-marker.tsv", sentence, "을"); got != particle {
			t.Errorf("%s: with no rule %s", sentence, got)
		}
		got, notes := ruledFirst(t, "words-open-marker.tsv", sentence, "을", openMarkerRule())
		if got != "을" {
			t.Errorf("%s: with the rule %s", sentence, got)
		}
		auxLinkedHas(t, notes, "  · "+particle+"\t"+ruleOpenMarker+"\t1.00\t통과")
	}
}
