//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMainUnlinkedRuleReader — read with the table of words-aux-linked.tsv: 있다 away from a linking form puts 있[열쇠:형용사] first with exactly bonusMainFirst more than it has with no rule while 있[열쇠:동사], which the rule does not speak of, keeps its preference; 보다 away from a linking form puts 보[열쇠:동사] first with bonusMainFirst more; after a linking form (먹어 보다) the main verb 보 is not raised
package parse

import "testing"

func TestMainUnlinkedRuleReader(t *testing.T) {
	for _, c := range []struct {
		in, main string
		raised   bool
	}{
		{"집에 있다", "있[열쇠:형용사]+다[맺음:다]", true},
		{"영화도 보다", "보[열쇠:동사]+다[맺음:다]", true},
		{"먹어 보다", "보[열쇠:동사]+다[맺음:다]", false},
		{"먹고 있다", "있[열쇠:형용사]+다[맺음:다]", false},
	} {
		got, _ := auxLinkedLast(t, c.in, true, false)
		bare, _ := auxLinkedLast(t, c.in, false, false)
		gain := auxLinkedPick(t, got, c.main).Pref - auxLinkedPick(t, bare, c.main).Pref
		if c.raised && !almost(gain, bonusMainFirst) || !c.raised && !almost(gain, 0) {
			t.Fatalf("%q: %v", c.in, gain)
		}
	}
	got, _ := auxLinkedLast(t, "집에 있다", true, false)
	bare, _ := auxLinkedLast(t, "집에 있다", false, false)
	verb := "있[열쇠:동사]+다[맺음:다]"
	if !almost(auxLinkedPick(t, got, verb).Pref, auxLinkedPick(t, bare, verb).Pref) {
		t.Fatal("있 the verb is raised")
	}
}
