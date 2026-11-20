//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestHaLinkedRuleStream — read as a stream with words-ha-linked.tsv (the numeral 하나 read before the verb 하 with the ending 나, by less than bonusShortTaken): with no rule 하나 is the numeral; with the rule 하나 after -기는 (있기는) and after -어야 (있어야) is 하 with the ending 나 and the trace says 통과, and 하나 after a noun with 은 stays the numeral and the trace says 기각: whyNoHaLink
package parse

import "testing"

func TestHaLinkedRuleStream(t *testing.T) {
	r := haLinkedRule()
	verb := "하[열쇠:동사]+나[맺음:나]"
	line := func(v, why string) string {
		return "  · " + verb + "\t" + ruleHaLinked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"돈이 있기는 하나", verb, line("1.00", "통과")},
		{"돈이 있어야 하나", verb, line("1.00", "통과")},
		{"돈은 하나", "하나[열쇠:체언]", line("-1.00", "기각: "+whyNoHaLink)},
	} {
		if got, _ := ruledFirst(t, "words-ha-linked.tsv", c.sentence, "하나"); got != "하나[열쇠:체언]" {
			t.Errorf("%s: 하나 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-ha-linked.tsv", c.sentence, "하나", r)
		if got != c.ruled {
			t.Errorf("%s: 하나 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
