//ff:func feature=parse type=rule control=sequence
//ff:what TestNominalBoundRuleStream — read as a stream with words-bound-table.tsv (the noun 측은 counted more than the dependent noun 측): with no rule 측은 after the noun 회사 is the noun 측은; with the rule it is 측+은 and the trace says 통과, and 측은 at the head of a sentence stays the noun and the trace says 기각: whyNotAfterNominal
package parse

import "testing"

func TestNominalBoundRuleStream(t *testing.T) {
	const table, apart = "words-bound-table.tsv", "측[열쇠:체언]+은[보조사:은]"
	r := nominalBoundRule(nominalBound)
	if got, _ := tableFirst(t, table, "회사 측은 말한다", "측은"); got != "측은[열쇠:체언]" {
		t.Fatalf("with no rule: %s", got)
	}
	got, notes := tableFirst(t, table, "회사 측은 말한다", "측은", r)
	if got != apart {
		t.Errorf("with the rule: %s, want %s", got, apart)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleNominalBound+"\t1.00\t통과")
	got, notes = tableFirst(t, table, "측은 말한다", "측은", r)
	if got != "측은[열쇠:체언]" {
		t.Errorf("at the head: %s", got)
	}
	auxLinkedHas(t, notes, "  · "+apart+"\t"+ruleNominalBound+"\t-1.00\t기각: "+whyNotAfterNominal)
}
