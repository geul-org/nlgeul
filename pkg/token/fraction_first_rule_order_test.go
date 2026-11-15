//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFractionFirstRuleOrder — over the small word table for streams (an index with a word table holds the rules of a sentence), with the rule numbers joined by a slash in the shape of a date (1/12, 3/20) are read first as a fraction, the date of the same letters right behind it, no candidate taken away, the two differing by exactly bonusOrder; with no rule (st.rules nil) the date is first, so it is the rule that changes the order; numbers that no date fits (34/1000) are a fraction alone as before
package token

import "testing"

func TestFractionFirstRuleOrder(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, text := range []string{"1/12", "3/20"} {
		part, date := text+"[나눔]", text+"[날짜]"
		bare := r.Stream(text)
		bare.rules = nil
		plain, _ := bare.Next()
		was := plain.Nodes[0]
		if was.Contest.Branches[0].String() != date || was.Contest.Branches[1].String() != part {
			t.Fatalf("%s without the rule: date first: %q", text, was.Contest.Branches)
		}
		seg, _ := r.Stream(text).Next()
		now := seg.Nodes[0]
		if len(now.Contest.Branches) != len(was.Contest.Branches) || now.Contest.Branches[0].String() != part || now.Contest.Branches[1].String() != date {
			t.Fatalf("%s: fraction first, date right behind: %q", text, now.Contest.Branches)
		}
		p, _ := branchNamed(now, part)
		d, _ := branchNamed(now, date)
		if !almost(p.Pref-d.Pref, bonusOrder) || p.Verdict != d.Verdict {
			t.Fatalf("%s: the difference is bonusOrder: %v %v", text, p.Pref, d.Pref)
		}
	}
	seg, _ := r.Stream("34/1000").Next()
	if got := seg.Nodes[0].Contest.Branches; got[0].String() != "34/1000[나눔]" {
		t.Fatal("no date: a fraction", got)
	}
}
