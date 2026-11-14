//ff:func feature=token type=reader control=sequence
//ff:what edgesOf — the signs at the edges of a unit that are not read with its letters: those before the letters and those after them, and where the latter begin in the unit (in letters); a unit with a sign of a pair or a span of code at its head has none before; a unit with nothing to read is all signs, given as those after, from 0
package token

func edgesOf(u unit, core []rune) ([]rune, []rune, int) {
	all := []rune(u.text)
	if u.head > 0 {
		return nil, all[len(core):], len(core)
	}
	from, to := trimmed(u.text)
	return all[:from], all[to:], to
}
