//ff:func feature=token type=helper control=sequence
//ff:what first — whether a state of the whole unit is a first word that has the sign of a pair at the unit's head as its first part (the candidates that took the sign into a word's letters are not)
package token

func (pl *planter) first(q *path) bool {
	parts := q.last.word.Parts
	return q.start == 0 && len(parts) > 0 && parts[0].Name.Kind == pairName
}
