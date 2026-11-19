//ff:func feature=token type=rule control=iteration dimension=1
//ff:what judgeCandidate — one candidate judged by each rule of a sentence that speaks of it (about): the graph is evaluated with the candidate and the words around it; when a rule's claim stands its bonus is gained and, for a rule of rebuttal, its force is kept; with a trace a line is written for each rule asked (the candidate, the rule's name, the verdict, 통과 or the reasons it fell); the sum gained, the forces of the rebuttals that stand and the lines are given back
package token

import (
	"fmt"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (s *Stream) judgeCandidate(br Branch, around rule.Around) (float64, []float64, []string) {
	add := 0.0
	var forces []float64
	var notes []string
	for _, r := range s.rules {
		if !r.about(around.Self) {
			continue
		}
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, around)
		v, why := evaluate(r.graph, ctx, s.trace)
		if s.trace {
			notes = append(notes, fmt.Sprintf("  · %s\t%s\t%.2f\t%s", br.String(), r.name, v, noteOf(v, why)))
		}
		if v > 0 {
			add += r.bonus
		}
		if v > 0 && r.force > 0 {
			forces = append(forces, r.force)
		}
	}
	return add, forces, notes
}
