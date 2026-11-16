//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSignShapesSlash — numbers joined by a slash in the shape of a date (1/12, 3/20) are read first as a fraction, by the order of the shapes of the slash (signShapes: 나눔 before 날짜), with the date of the same letters right behind it, of the same verdict and preference, with the rules of a sentence or without them; numbers that no date fits (34/1000) are a fraction alone
package token

import "testing"

func TestSignShapesSlash(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, c := range []struct {
		text  string
		ruled bool
	}{{"1/12", true}, {"1/12", false}, {"3/20", true}, {"3/20", false}} {
		part, date := c.text+"[나눔]", c.text+"[날짜]"
		st := r.Stream(c.text)
		if !c.ruled {
			st.rules = nil
		}
		seg, _ := st.Next()
		now := seg.Nodes[0]
		if now.Contest.Branches[0].String() != part || now.Contest.Branches[1].String() != date {
			t.Fatalf("%s (rules %v): fraction first, date right behind: %q", c.text, c.ruled, now.Contest.Branches)
		}
		p, _ := branchNamed(now, part)
		d, _ := branchNamed(now, date)
		if p.Pref != d.Pref || p.Verdict != d.Verdict {
			t.Fatalf("%s: the same preference and verdict: %v %v", c.text, p, d)
		}
	}
	seg, _ := r.Stream("34/1000").Next()
	if got := seg.Nodes[0].Contest.Branches; got[0].String() != "34/1000[나눔]" {
		t.Fatal("no date: a fraction", got)
	}
}
