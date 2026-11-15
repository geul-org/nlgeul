//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stretchFirst — the first line of the stretch of a sentence whose letters are the given ones, as the stream sends it first (Next: the first candidate of each node of its best way, joined with " "); "" when no stretch has those letters
package token

func stretchFirst(r *Reader, sentence, letters string) string {
	for _, sg := range streamFirsts(r.Stream(sentence)) {
		if first := firstsOf(sg, letters); len(first) > 0 {
			return first[0]
		}
	}
	return ""
}
