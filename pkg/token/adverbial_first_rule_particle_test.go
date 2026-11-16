//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAdverbialFirstRuleParticle — with a particle after 에서 (학교에서는) the adverbial reading, which has a marker 에서 followed by 는, is still the one the rule raises: it is first, by bonusSameLetters over what it was with no rule, and every other candidate keeps its preference
package token

import "testing"

func TestAdverbialFirstRuleParticle(t *testing.T) {
	const adv = "학교[열쇠:체언]+에서[부사격:에서]+는[보조사:는]"
	r := streamReader(t, "words-adverbial.tsv")
	bare := r.Stream("학교에서는")
	bare.rules = nil
	plain, _ := bare.Next()
	seg, _ := r.Stream("학교에서는").Next()
	was, now := plain.Nodes[0], seg.Nodes[0]
	if now.Contest.Branches[0].String() != adv || len(now.Contest.Branches) != len(was.Contest.Branches) {
		t.Fatal(now.Contest.Branches)
	}
	for _, br := range now.Contest.Branches {
		old, i := branchNamed(was, br.String())
		if i < 0 {
			t.Fatal("new candidate", br.String())
		}
		want := old.Pref
		if br.String() == adv || br.String() == "학교에+서[부사격:서]+는[보조사:는]" {
			want += bonusSameLetters
		}
		if !almost(br.Pref, want) {
			t.Errorf("%s: %v; want %v", br.String(), br.Pref, want)
		}
	}
}
