//ff:func feature=token type=rule control=sequence
//ff:what TestBoundTableRuleStream — read as a stream with words-bound-table.tsv (the noun 리가 counted more than the dependent noun 리): with no rule 리가 after the modifying ㄹ is the noun 리가; with the rule it is 리+가 and the trace says 통과, and 리가 at the head of a sentence stays the noun and the trace says 기각: whyNoAdn
package token

import "testing"

func TestBoundTableRuleStream(t *testing.T) {
	const table, apart = "words-bound-table.tsv", "리[열쇠:체언]+가[주격:가]"
	r := boundTableRule(tableBound)
	if got, _ := tableFirst(t, table, "들릴 리가 없다", "리가"); got != "리가[열쇠:체언]" {
		t.Fatalf("with no rule: %s", got)
	}
	got, notes := tableFirst(t, table, "들릴 리가 없다", "리가", r)
	if got != apart {
		t.Errorf("with the rule: %s, want %s", got, apart)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleBoundTable+"\t1.00\t통과")
	got, notes = tableFirst(t, table, "리가 없다", "리가", r)
	if got != "리가[열쇠:체언]" {
		t.Errorf("at the head: %s", got)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleBoundTable+"\t-1.00\t기각: "+whyNoAdn)
}
