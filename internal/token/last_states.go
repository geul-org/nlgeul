//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lastStates — the states a stretch's best way ends in: the states at its first span, then at each span the states that have one at the span before among theirs; those at the last span are the states any way over the same spans can go on from (a state of the next stretch stands on one of them); none when no state is left
package token

func lastStates(pt portion, seg Segment) []*path {
	var set []*path
	for i, nd := range seg.Nodes {
		set = statesAfter(pt.states, nd.Span, set, i == 0)
	}
	return set
}
