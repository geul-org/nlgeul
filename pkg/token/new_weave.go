//ff:func feature=token type=helper control=iteration dimension=1
//ff:what newWeave — a portion of the net laid out: the spans where its states stand sorted by start and then end (the order the old cuts give), which states stand at each, the states before each that lie in the portion, and the spans that follow each span
package token

import "sort"

func newWeave(ix *index, pt portion) *weave {
	w := &weave{ix: ix, from: pt.from, to: pt.to, states: pt.states, at: map[Span]int{}, mark: make([]int, len(pt.states))}
	for _, p := range pt.states {
		sp := Span{Start: p.start, End: p.last.end}
		if _, ok := w.at[sp]; !ok {
			w.at[sp] = 0
			w.spans = append(w.spans, sp)
		}
	}
	sort.Slice(w.spans, func(i, j int) bool { return spanBefore(w.spans[i], w.spans[j]) })
	for i, sp := range w.spans {
		w.at[sp] = i
	}
	w.of = make([][]int, len(w.spans))
	idOf := map[*path]int{}
	for id, p := range pt.states {
		idOf[p] = id
		w.node = append(w.node, w.at[Span{Start: p.start, End: p.last.end}])
		w.of[w.node[id]] = append(w.of[w.node[id]], id)
	}
	w.link(idOf)
	return w
}
