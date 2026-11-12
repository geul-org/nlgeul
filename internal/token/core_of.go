//ff:func feature=token type=reader control=sequence
//ff:what coreOf — the letters of a unit that are read: a plain unit without the signs at its edges (trimUnit); a unit with a sign of a pair or a span of code at its head keeps the head, and what is written onto it loses the signs at its edges
package token

func coreOf(u unit) []rune {
	if u.head == 0 {
		return trimUnit(u.text)
	}
	r := []rune(u.text)
	return append(append([]rune{}, r[:u.head]...), trimSigns(string(r[u.head:]))...)
}
