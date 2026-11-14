//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sameRuns — some signs cut where the sign changes: each stretch of one sign in a row, as where it starts and ends
package token

func sameRuns(signs []rune) [][2]int {
	var out [][2]int
	start := 0
	for i := 1; i <= len(signs); i++ {
		if i == len(signs) || signs[i] != signs[start] {
			out, start = append(out, [2]int{start, i}), i
		}
	}
	return out
}
