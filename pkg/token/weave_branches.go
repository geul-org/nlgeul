//ff:func feature=token type=helper control=iteration dimension=1
//ff:what branches — the candidates at a span, one for each way a word is printed there, in the order the words were made (branchesOf for the states that stand at the span)
package token

func (w *weave) branches(node int) []Branch {
	states := make([]*path, 0, len(w.of[node]))
	for _, id := range w.of[node] {
		states = append(states, w.states[id])
	}
	return branchesOf(w.ix, states)
}
