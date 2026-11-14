//ff:func feature=token type=helper control=iteration dimension=2
//ff:what covered — the places of a sentence that lie inside one of some pairs (from the opening sign to the closing one)
package token

func covered(pairs map[int]int, n int) []bool {
	out := make([]bool, n)
	for open, shut := range pairs {
		for i := open; i <= shut; i++ {
			out[i] = true
		}
	}
	return out
}
