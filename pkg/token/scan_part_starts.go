//ff:func feature=token type=helper control=iteration dimension=1
//ff:what partStarts — when the unit holds a slash or a hyphen (the signs words are bound by in one stem), for each place where the run of letters with neither sign that ends there begins; nothing for a unit with no such sign
package token

import "slices"

func (sc *scan) partStarts() {
	sc.parts = nil
	if !slices.Contains(sc.s, '/') && !slices.Contains(sc.s, '-') {
		return
	}
	sc.parts = make([]int, len(sc.s)+1)
	for p := 1; p <= len(sc.s); p++ {
		sc.parts[p] = sc.parts[p-1]
		if sc.s[p-1] == '/' || sc.s[p-1] == '-' {
			sc.parts[p] = p
		}
	}
}
