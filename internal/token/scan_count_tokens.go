//ff:func feature=token type=helper control=iteration dimension=1
//ff:what countTokens — for each place how many tokens lie before it (an atom is one; the places inside an atom are not counted), so that the tokens of a stem are the difference of two numbers
package token

func (sc *scan) countTokens() {
	sc.count = make([]int, len(sc.s)+1)
	for p := 1; p <= len(sc.s); p++ {
		sc.count[p] = sc.count[p-1]
		if !sc.inside[p] {
			sc.count[p]++
		}
	}
}
