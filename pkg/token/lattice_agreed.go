//ff:func feature=token type=helper control=iteration dimension=1
//ff:what agreed — the places where every way of reading the unit puts a boundary: a place from 0 to the unit's end that no live state (a state on a way that reads the whole unit) lies across, with the place the letters begin at (0) first and the unit's end last; none at all when no way reads the unit
package token

func (l *lattice) agreed(live []*path) []int {
	if len(live) == 0 {
		return nil
	}
	across := make([]int, l.n+2)
	for _, p := range live {
		if p.last.end-p.start >= 2 {
			across[p.start+1]++
			across[p.last.end]--
		}
	}
	out := []int{l.from}
	run := 0
	for q := l.from + 1; q < l.n; q++ {
		run += across[q]
		if run == 0 {
			out = append(out, q)
		}
	}
	return append(out, l.n)
}
