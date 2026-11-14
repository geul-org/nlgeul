//ff:func feature=token type=reader control=iteration dimension=1
//ff:what newPlanter — the whole unit of a sign of a pair with letters written onto it, ready to be read place by place: only the words that begin at the unit's head or after the sign are made (starts), and the places up to the sign are worked out
package token

func newPlanter(sc *scan, head int) *planter {
	sc.starts = map[int]bool{0: true, head: true}
	pl := &planter{sc: sc, best: sc.netStart(), head: head, transparent: map[*path]bool{}}
	pl.signPref = sc.wordOf(0, head, tail{start: head}).pref
	for j := 1; j <= head; j++ {
		pl.step(j)
	}
	return pl
}
