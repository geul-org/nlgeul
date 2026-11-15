//ff:func feature=token type=rule control=iteration dimension=1
//ff:what judgeSeat — the i-th word of the line judged by the rules of a sentence: each of its candidates with the words before (nearsBack) and after (nearsOn) it; what the candidates gain (lifts) and the trace lines are given back, and the seat reads as its first candidate after the lifts
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (s *Stream) judgeSeat(seats []seat, i int) ([]lift, []string) {
	n := seats[i].node
	around := rule.Around{Before: nearsBack(seats[:i], s.behind), After: nearsOn(seats[i+1:])}
	var lifts []lift
	var notes []string
	for _, br := range n.Contest.Branches {
		around.Self = nearIn(n, br)
		around.Self.Text = n.Contest.Text
		add, lines := s.judgeCandidate(br, around)
		notes = append(notes, lines...)
		if add != 0 {
			lifts = append(lifts, lift{span: n.Span, word: br.String(), add: add})
		}
	}
	seats[i].near = nearIn(n, lifted(n, lifts).Contest.Branches[0])
	return lifts, notes
}
