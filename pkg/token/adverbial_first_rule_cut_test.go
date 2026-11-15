//ff:func feature=token type=reader control=sequence
//ff:what TestAdverbialFirstRuleCut — the rule does not carry a candidate over another cut: with 계획서 and 계획 both in the table 계획서는 stays 계획서+는 first (계획+서+는, which the rule raises, is still a candidate behind it), and with the same rule given bonusSentence in place of bonusOrder the order is the other way, so the small value is what holds the cut
package token

import "testing"

func TestAdverbialFirstRuleCut(t *testing.T) {
	const whole, split = "계획서[열쇠:체언]+는[보조사:는]", "계획[열쇠:체언]+서[부사격:서]+는[보조사:는]"
	r := streamReader(t, "words-adverbial.tsv")
	bare := r.Stream("계획서는")
	bare.rules = nil
	plain, _ := bare.Next()
	was := plain.Nodes[0]
	if was.Contest.Branches[0].String() != whole {
		t.Fatalf("the table must put the whole noun first with no rule: %q", was.Contest.Branches)
	}
	seg, _ := r.Stream("계획서는").Next()
	now := seg.Nodes[0]
	w, wi := branchNamed(now, whole)
	s, si := branchNamed(now, split)
	if wi != 0 || si < 0 {
		t.Fatalf("whole first, split kept: %q", now.Contest.Branches)
	}
	oldSplit, _ := branchNamed(was, split)
	if !almost(s.Pref, oldSplit.Pref+bonusOrder) || w.Lambda() <= s.Lambda() {
		t.Fatalf("lambda %v %v", w.Lambda(), s.Lambda())
	}

	loud := adverbialFirstRule()
	loud.bonus = bonusSentence
	st := r.Stream("계획서는")
	st.rules = []sentenceRule{loud}
	seg, _ = st.Next()
	if got := seg.Nodes[0].Contest.Branches[0].String(); got != split {
		t.Fatalf("with bonusSentence the cut moves to the split: %q", got)
	}
}
