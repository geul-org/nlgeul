//ff:func feature=token type=reader control=iteration dimension=1
//ff:what twoRuns — do the letters from b to j read as two runs of markers one after the other, each as the index reads it where it ends (는 and then 게 in 한다 는게): such letters are markers though no one chain holds them
package token

func (sc *scan) twoRuns(b, j int) bool {
	for k := b + 1; k < j; k++ {
		if sc.marks[k][b] && sc.marks[j][k] {
			return true
		}
	}
	return false
}
