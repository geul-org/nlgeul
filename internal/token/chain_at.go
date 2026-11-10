//ff:func feature=token type=reader control=iteration dimension=1
//ff:what chainAt — the name or number that begins with the k-th run: that run and the runs after it that a single dot ties on, or a single sign of numbers between two runs of digits
package token

func chainAt(s []rune, runs [][2]int, k int) [][2]int {
	out := [][2]int{runs[k]}
	for k+1 < len(runs) && runs[k+1][0] == runs[k][1]+1 && ties(s, runs[k], runs[k+1]) {
		k++
		out = append(out, runs[k])
	}
	return out
}
