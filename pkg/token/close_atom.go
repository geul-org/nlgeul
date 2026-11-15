//ff:func feature=token type=reader control=sequence
//ff:what closeAtom — note one atom: a number is named so and takes the sign at its very front (a minus or plus, money) when that sign's graph stands, or else the prefix 제 right before it when its graph stands; the places inside are marked and a number's end is kept, and noted too when the number begins with 제
package token

func (sc *scan) closeAtom(start, end int, names []string, num bool) {
	if num && len(names) == 0 {
		names = []string{numName}
	}
	if num {
		was := start
		start, names = sc.front(start, names)
		if start == was {
			start, names = sc.orderFront(start, names)
			sc.ordEnds[end] = start < was
		}
		sc.numEnds[end] = true
	}
	sc.atoms[start] = atom{end, names, num}
	sc.fill(start, end)
}
