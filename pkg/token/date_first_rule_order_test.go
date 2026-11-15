//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestDateFirstRuleOrder — over the small word table for streams (an index with a word table holds the rules of a sentence), with the rule a number of full stops in the shape of a date (1989.8.25, 1964.7.13) is read first as a date, the version of the same letters right behind it, no candidate taken away, the two differing by exactly bonusOrder; with no rule (st.rules nil) the version is first, so it is the rule that changes the order; numbers of full stops that are no date (1.2.3) are read as before
package token

import "testing"

func TestDateFirstRuleOrder(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, text := range []string{"1989.8.25", "1964.7.13"} {
		date, version := text+"[날짜]", text+"[버전]"
		bare := r.Stream(text)
		bare.rules = nil
		plain, _ := bare.Next()
		was := plain.Nodes[0]
		if was.Contest.Branches[0].String() != version || was.Contest.Branches[1].String() != date {
			t.Fatalf("%s without the rule: version first: %q", text, was.Contest.Branches)
		}
		seg, _ := r.Stream(text).Next()
		now := seg.Nodes[0]
		if len(now.Contest.Branches) != len(was.Contest.Branches) || now.Contest.Branches[0].String() != date || now.Contest.Branches[1].String() != version {
			t.Fatalf("%s: date first, version right behind: %q", text, now.Contest.Branches)
		}
		d, _ := branchNamed(now, date)
		v, _ := branchNamed(now, version)
		if !almost(d.Pref-v.Pref, bonusOrder) || d.Verdict != v.Verdict {
			t.Fatalf("%s: the difference is bonusOrder: %v %v", text, d.Pref, v.Pref)
		}
	}
	seg, _ := r.Stream("1.2.3").Next()
	if got := seg.Nodes[0].Contest.Branches[0].String(); got != "1.2.3[버전]" {
		t.Fatal("no date: the version stays first", got)
	}
}
