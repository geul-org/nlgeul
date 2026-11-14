//ff:func feature=token type=helper control=iteration dimension=2
//ff:what branchLines — the branches of all the contests of a sentence, each printed (Branch.String), as a set
package token

func branchLines(r *Reader, sentence string) map[string]bool {
	out := map[string]bool{}
	for _, c := range r.Contests(sentence) {
		for _, b := range c.Branches {
			out[b.String()] = true
		}
	}
	return out
}
