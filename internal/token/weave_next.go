//ff:func feature=token type=helper control=iteration dimension=2
//ff:what next — the states at a span that have a state of the set before them: the set of states that a way has come to when it goes on to that span (empty when no state there follows one of the set); the set and the result are in the order the states were made
package token

func (w *weave) next(set []int, node int) []int {
	w.tick++
	for _, id := range set {
		w.mark[id] = w.tick
	}
	var out []int
	for _, id := range w.of[node] {
		for _, r := range w.preds[id] {
			if w.mark[r] == w.tick {
				out = append(out, id)
				break
			}
		}
	}
	return out
}
