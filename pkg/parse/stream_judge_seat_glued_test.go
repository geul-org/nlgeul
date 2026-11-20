//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestStreamJudgeSeatGlued — each candidate of a seat is judged as its own word with whether no space stands before it and the word's letters (Glued and Text of the seat's node): 학생밥을 먹었다 has 밥을 glued to 학생 with its letters, 학생 and 먹었다 not glued
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStreamJudgeSeatGlued(t *testing.T) {
	var seen []rule.Near
	look := sentenceRule{name: "본 것", graph: toulmin.NewGraph("본 것"), about: func(n rule.Near) bool {
		seen = append(seen, n)
		return false
	}}
	st := sentenceStream(t, "", false, look)
	seats := seatsOf(sentenceRegions(t, "학생밥을 먹었다"))
	var texts []string
	for i := range seats {
		seen = nil
		st.judgeSeat(seats, i)
		n := seats[i].node
		texts = append(texts, n.Contest.Text)
		if len(seen) != len(n.Contest.Branches) {
			t.Fatal(i, len(seen), len(n.Contest.Branches))
		}
		if slices.ContainsFunc(seen, func(s rule.Near) bool { return s.Glued != n.Contest.Glued || s.Text != n.Contest.Text }) {
			t.Errorf("seat %d: %+v, node glued %v %q", i, seen, n.Contest.Glued, n.Contest.Text)
		}
	}
	if len(seats) != 3 || texts[1] != "밥을" || !seats[1].node.Contest.Glued || seats[0].node.Contest.Glued || seats[2].node.Contest.Glued {
		t.Fatal(texts)
	}
}
