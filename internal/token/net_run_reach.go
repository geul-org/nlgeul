//ff:func feature=token type=reader control=iteration dimension=1
//ff:what reach — the best way to every state kept, from the states before the first (the seeds, or the root): the states in the order of where they end (every state before one is before it), each given the best value of the states before it plus the best lambda of its span (the best of the states at the span as the stretch of that span would give: spanLambda); the first state before it to come wins a tie
package token

func (r *netRun) reach() *reach {
	order := r.byEnd()
	rc := &reach{value: map[*path]float64{}, back: map[*path]*path{}, group: map[Span][]*path{}}
	for _, s := range order {
		sp := Span{Start: s.start, End: s.last.end}
		rc.group[sp] = append(rc.group[sp], s)
	}
	for _, s := range order {
		r.score(rc, s)
	}
	return rc
}
