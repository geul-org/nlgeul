//ff:func feature=token type=reader control=sequence
//ff:what closeAtom — note one atom: a number is named so and takes the sign at its very front (a minus or plus, money) when that sign's graph stands; the places inside are marked and a number's end is kept
package token

func (sc *scan) closeAtom(start, end int, names string, num bool) {
	if num && names == "" {
		names = numName
	}
	if num {
		start, names = sc.front(start, names)
		sc.numEnds[end] = true
	}
	sc.atoms[start] = atom{end, names, num}
	sc.fill(start, end)
}
