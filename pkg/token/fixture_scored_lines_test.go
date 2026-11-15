//ff:func feature=token type=helper control=iteration dimension=2
//ff:what scoredLines — streamLines with the values of each line after a tab (wayScored: the least verdict and the summed preference of the chosen candidates)
package token

func scoredLines(r *Reader, sentence string) []string {
	var out []string
	st := r.Stream(sentence)
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		for _, way := range streamWays(st, seg) {
			out = append(out, wayScored(way)...)
		}
	}
	return out
}
