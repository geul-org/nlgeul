//ff:func feature=token type=helper control=iteration dimension=1
//ff:what member — is a line of spans a way of cutting the portion: it begins where the portion begins and ends where it ends, each span is one where words stand, and the words can follow one another (a word at each span that has a word of the span before it among the states before it, so that the line is a chain of words and not only of spans)
package token

func (w *weave) member(spans []Span) bool {
	if len(spans) == 0 || spans[0].Start != w.from || spans[len(spans)-1].End != w.to {
		return false
	}
	node, ok := w.at[spans[0]]
	if !ok {
		return false
	}
	set := w.of[node]
	for _, sp := range spans[1:] {
		node, ok = w.at[sp]
		if !ok {
			return false
		}
		if set = w.next(set, node); len(set) == 0 {
			return false
		}
	}
	return true
}
