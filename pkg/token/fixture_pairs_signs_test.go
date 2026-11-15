//ff:func feature=token type=helper control=iteration dimension=1
//ff:what pairsSigns — the contests of a sentence whose text is the given sign, in order
package token

func pairsSigns(r *Reader, sentence, sign string) []Contest {
	var out []Contest
	for _, b := range streamContests(r, sentence) {
		if b.Text == sign {
			out = append(out, b)
		}
	}
	return out
}
