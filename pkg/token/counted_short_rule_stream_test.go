//ff:func feature=token type=rule control=sequence
//ff:what TestCountedShortRuleStream — read as a stream with words-bound-table.tsv (the noun 대가 counted more than the noun 대 with 가): with no rule 대가 after the numeral determiner 한 is the noun 대가; with the rule it is 대+가 and the trace says 통과, and 대가 after the demonstrative 그 stays the noun and the trace says 기각
package token

import "testing"

func TestCountedShortRuleStream(t *testing.T) {
	const table, apart = "words-bound-table.tsv", "대[열쇠:체언]+가[주격:가]"
	r := countedShortRule(numDeterminers)
	if got, _ := tableFirst(t, table, "차 한 대가 지나간다", "대가"); got != "대가[열쇠:체언]" {
		t.Fatalf("with no rule: %s", got)
	}
	got, notes := tableFirst(t, table, "차 한 대가 지나간다", "대가", r)
	if got != apart {
		t.Errorf("with the rule: %s, want %s", got, apart)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleCountedShort+"\t1.00\t통과")
	got, notes = tableFirst(t, table, "차 그 대가 지나간다", "대가", r)
	if got != "대가[열쇠:체언]" {
		t.Errorf("after 그: %s", got)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleCountedShort+"\t-1.00\t기각: "+whyNoNumDet)
}
