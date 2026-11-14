//ff:func feature=token type=reader control=sequence
//ff:what trimUnit — a unit without the signs at its edges, keeping a sign that belongs to a number there (one that stands at a number's front: -3, $40; a sign of the unit table at its end: 50%, 3℃) and the signs that belong to a name: a path's head and the slash at its end, the dashes of an option, and $ _ . before Latin letters (trimmed)
package token

func trimUnit(field string) []rune {
	from, to := trimmed(field)
	return []rune(field)[from:to]
}
