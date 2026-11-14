//ff:func feature=token type=reader control=iteration dimension=1
//ff:what lines — the spans along the best way to each of the states given, first span first (reach and back from the state to a seed or the root)
package token

import "slices"

func (r *netRun) lines(front []*path) [][]Span {
	rc := r.reach()
	out := make([][]Span, 0, len(front))
	for _, f := range front {
		var line []Span
		for p := f; p != nil && p.prev != nil; p = rc.back[p] {
			line = append(line, Span{Start: p.start, End: p.last.end})
		}
		slices.Reverse(line)
		out = append(out, line)
	}
	return out
}
