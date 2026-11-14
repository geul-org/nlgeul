//ff:func feature=token type=reader control=sequence
//ff:what stand — make the letters from b up to j a word that may stand as a word of no markers inside the unit (stands)
package token

func (sc *scan) stand(j, b int) {
	if sc.stands[j] == nil {
		sc.stands[j] = map[int]bool{}
	}
	sc.stands[j][b] = true
}
