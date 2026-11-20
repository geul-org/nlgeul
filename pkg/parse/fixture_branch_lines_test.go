//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what branchLines — what the stream of a sentence gives as lines (streamLines: every way of every stretch, one candidate of each word), as a set
package parse

func branchLines(r *Reader, sentence string) map[string]bool {
	out := map[string]bool{}
	for _, l := range streamLines(r, sentence) {
		out[l] = true
	}
	return out
}
