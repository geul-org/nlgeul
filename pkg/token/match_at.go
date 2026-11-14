//ff:func feature=token type=helper control=iteration dimension=1
//ff:what matchAt — where the word w ends if it stands in the letters s from place p (-1 when it does not, or when it runs past the end)
package token

// The letters are compared one by one, with no copy of either made.
func matchAt(s []rune, p int, w string) int {
	q := p
	for _, r := range w {
		if q >= len(s) || s[q] != r {
			return -1
		}
		q++
	}
	return q
}
