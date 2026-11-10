//ff:func feature=token type=reader control=iteration dimension=1
//ff:what fill — mark the places strictly inside an atom: no word begins or ends there
package token

func (sc *scan) fill(start, end int) {
	for p := start + 1; p < end; p++ {
		sc.inside[p] = true
	}
}
