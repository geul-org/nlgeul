//ff:func feature=token type=helper control=iteration dimension=1
//ff:what afterOf — a word of the whole unit that begins right after the sign alone (a state that begins at the head and has a state before it that is the sign alone), as it is carried over to the letters after the sign: its end and its word's end counted from there, all else as it was made
package token

func (pl *planter) afterOf(q *path) (step, bool) {
	if q.start != pl.head {
		return step{}, false
	}
	for _, p := range q.preds {
		if pl.transparent[p] {
			e := q.last
			e.end -= pl.head
			e.word.End = e.end
			return e, true
		}
	}
	return step{}, false
}
