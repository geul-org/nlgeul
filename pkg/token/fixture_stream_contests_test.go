//ff:func feature=token type=helper control=iteration dimension=1
//ff:what streamContests — what the stream of a sentence gives, set as the old structure for the tests that read it so: a contest for each stretch Next sends, in order (stretchContest: its letters, the unit and Glued of its first word, every way's choices of candidates as branches, the best way's first)
package token

func streamContests(r *Reader, sentence string) []Contest {
	var out []Contest
	letters := lettersOf(sentence)
	st := r.Stream(sentence)
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		out = append(out, stretchContest(st, seg, letters))
	}
	return out
}
