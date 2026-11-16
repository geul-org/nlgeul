//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddNumbersDot — numbers of full stops in the shape of a date (1989.8.25, 1964.7.13) are read first as a date, by the order of the candidates of the full stop (addNumbers: the shapes of a number before the joins), with the version of the same letters right behind it, of the same verdict and preference, with the rules of a sentence or without them; numbers of full stops that are no date (1.2.3) are read as a version first
package token

import "testing"

func TestAddNumbersDot(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, c := range []struct {
		text  string
		ruled bool
	}{{"1989.8.25", true}, {"1989.8.25", false}, {"1964.7.13", true}, {"1964.7.13", false}} {
		date, version := c.text+"[날짜]", c.text+"[버전]"
		st := r.Stream(c.text)
		if !c.ruled {
			st.rules = nil
		}
		seg, _ := st.Next()
		now := seg.Nodes[0]
		if now.Contest.Branches[0].String() != date || now.Contest.Branches[1].String() != version {
			t.Fatalf("%s (rules %v): date first, version right behind: %q", c.text, c.ruled, now.Contest.Branches)
		}
		d, _ := branchNamed(now, date)
		v, _ := branchNamed(now, version)
		if d.Pref != v.Pref || d.Verdict != v.Verdict {
			t.Fatalf("%s: the same preference and verdict: %v %v", c.text, d, v)
		}
	}
	seg, _ := r.Stream("1.2.3").Next()
	if got := seg.Nodes[0].Contest.Branches[0].String(); got != "1.2.3[버전]" {
		t.Fatal("no date: the version stays first", got)
	}
}
