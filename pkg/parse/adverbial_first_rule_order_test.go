//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAdverbialFirstRuleOrder — with the rule 에서 and 서 as the adverbial particle come before the same letters as the nominative (학교에서, 여기서; 학교에서는 has no nominative to pass), the nominative stays right behind, no candidate is taken away, the two differ by exactly bonusSameLetters in preference and the rest are as they were; with no rule (st.rules nil) the nominative is first, so it is the rule that changes the order
package parse

import "testing"

func TestAdverbialFirstRuleOrder(t *testing.T) {
	r := streamReader(t, "words-adverbial.tsv")
	cases := []struct{ text, adv, nom string }{
		{"학교에서", "학교[열쇠:체언]+에서[부사격:에서]", "학교[열쇠:체언]+에서[주격:에서]"},
		{"여기서", "여기[열쇠:체언]+서[부사격:서]", "여기[열쇠:체언]+서[주격:서]"},
	}
	for _, c := range cases {
		bare := r.Stream(c.text)
		bare.rules = nil
		plain, _ := bare.Next()
		was := plain.Nodes[0]
		if was.Contest.Branches[0].String() != c.nom || was.Contest.Branches[1].String() != c.adv {
			t.Fatalf("%s without the rule: nominative first: %q", c.text, was.Contest.Branches)
		}
		seg, ok := r.Stream(c.text).Next()
		if !ok {
			t.Fatal(c.text)
		}
		now := seg.Nodes[0]
		if len(now.Contest.Branches) != len(was.Contest.Branches) {
			t.Fatalf("%s: a candidate was taken away or added: %q", c.text, now.Contest.Branches)
		}
		if now.Contest.Branches[0].String() != c.adv || now.Contest.Branches[1].String() != c.nom {
			t.Fatalf("%s: adverbial first, nominative right behind: %q", c.text, now.Contest.Branches)
		}
		adv, _ := branchNamed(now, c.adv)
		nom, _ := branchNamed(now, c.nom)
		if !almost(adv.Pref-nom.Pref, bonusSameLetters) || adv.Verdict != nom.Verdict {
			t.Fatalf("%s: the difference is bonusSameLetters: %v %v", c.text, adv.Pref, nom.Pref)
		}
		oldAdv, _ := branchNamed(was, c.adv)
		oldNom, _ := branchNamed(was, c.nom)
		if !almost(adv.Pref, oldAdv.Pref+bonusSameLetters) || nom.Pref != oldNom.Pref {
			t.Fatalf("%s: only the adverbial is raised, by bonusSameLetters", c.text)
		}
	}
}
