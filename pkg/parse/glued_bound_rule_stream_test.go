//ff:func feature=parse type=rule control=sequence
//ff:what TestGluedBoundRuleStream — read as a stream with words-bound-table.tsv (the noun 자가 counted more than the dependent noun 자): in 패소자가 the noun 자가 comes first with no rule and with gluedShortRule alone; with gluedShortRule and gluedBoundRule 자+가 comes first and the trace says both stood
package parse

import "testing"

func TestGluedBoundRuleStream(t *testing.T) {
	const table, sentence, apart = "words-bound-table.tsv", "패소자가 부담한다", "자[열쇠:체언]+가[주격:가]"
	if got, _ := tableFirst(t, table, sentence, "자가"); got != "자가[열쇠:체언]" {
		t.Fatalf("with no rule: %s", got)
	}
	if got, _ := tableFirst(t, table, sentence, "자가", gluedShortRule()); got != "자가[열쇠:체언]" {
		t.Fatalf("with gluedShortRule alone: %s", got)
	}
	got, notes := tableFirst(t, table, sentence, "자가", gluedShortRule(), gluedBoundRule(tableBound))
	if got != apart {
		t.Errorf("with both rules: %s, want %s", got, apart)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleGluedShort+"\t1.00\t통과", "  · "+apart+"\t"+ruleGluedBound+"\t1.00\t통과")
}
