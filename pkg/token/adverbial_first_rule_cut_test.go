//ff:func feature=token type=reader control=sequence
//ff:what TestAdverbialFirstRuleCut — the rule carries a candidate over another reading of the word just above it: with 가 both a noun and a verb in the table, 가서는 has 가+(아)서+는 of the verb first with no rule, and with the rule 가+서+는, raised by exactly bonusSameLetters, comes first with the verb's still a candidate behind it
package token

import "testing"

func TestAdverbialFirstRuleCut(t *testing.T) {
	const verb, split = "가[열쇠:동사]+(아)서[맺음:어서(합침)]+는[보조사:는]", "가[열쇠:체언]+서[부사격:서]+는[보조사:는]"
	r := streamReader(t, "words-adverbial.tsv")
	bare := r.Stream("가서는")
	bare.rules = nil
	plain, _ := bare.Next()
	was := plain.Nodes[0]
	if was.Contest.Branches[0].String() != verb {
		t.Fatalf("the table must put the verb first with no rule: %q", was.Contest.Branches)
	}
	seg, _ := r.Stream("가서는").Next()
	now := seg.Nodes[0]
	w, wi := branchNamed(now, verb)
	s, si := branchNamed(now, split)
	if si != 0 || wi < 0 {
		t.Fatalf("split first, verb kept: %q", now.Contest.Branches)
	}
	oldSplit, _ := branchNamed(was, split)
	if !almost(s.Pref, oldSplit.Pref+bonusSameLetters) || w.Lambda() >= s.Lambda() {
		t.Fatalf("lambda %v %v", w.Lambda(), s.Lambda())
	}
}
