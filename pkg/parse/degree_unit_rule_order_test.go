//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestDegreeUnitRuleOrder — over the small word table for streams (an index with a word table holds the rules of a sentence), with the rule 도 right after a number at the end of its word (10도, 340도) is read first as the unit, the particle of the same letters behind it, no candidate taken away, the unit raised by exactly bonusSameLetters; with no rule (st.rules nil) the particle is first, so it is the rule that changes the order; 10도로 (a particle after the unit) is read as before
package parse

import "testing"

func TestDegreeUnitRuleOrder(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, text := range []string{"10도", "340도"} {
		n := text[:len(text)-len("도")]
		unit, particle := n+"[수:십진]+도[단위:도]", n+"[수:십진]+도[보조사:도]"
		bare := r.Stream(text)
		bare.rules = nil
		plain, _ := bare.Next()
		was := plain.Nodes[0]
		if was.Contest.Branches[0].String() != particle {
			t.Fatalf("%s without the rule: the particle first: %q", text, was.Contest.Branches)
		}
		seg, _ := r.Stream(text).Next()
		now := seg.Nodes[0]
		if len(now.Contest.Branches) != len(was.Contest.Branches) || now.Contest.Branches[0].String() != unit {
			t.Fatalf("%s: the unit first: %q", text, now.Contest.Branches)
		}
		u, _ := branchNamed(now, unit)
		old, _ := branchNamed(was, unit)
		if !almost(u.Pref, old.Pref+bonusSameLetters) {
			t.Fatalf("%s: raised by bonusSameLetters: %v %v", text, u.Pref, old.Pref)
		}
	}
	seg, _ := r.Stream("10도로").Next()
	if got := seg.Nodes[0].Contest.Branches[0].String(); got != "10[수:십진]+도[단위:도]+로[부사격:로]" {
		t.Fatal("a particle after the unit", got)
	}
}
