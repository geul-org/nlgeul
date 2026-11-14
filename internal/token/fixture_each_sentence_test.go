//ff:func feature=token type=helper control=iteration dimension=1
//ff:what eachSentence — the sentences the stream tests read, each put to a check with the reader; what the check names that is wrong, each with its sentence
package token

func eachSentence(r *Reader, check func(*Reader, string) []string) []string {
	var bad []string
	for _, s := range streamSentences() {
		for _, b := range check(r, s) {
			bad = append(bad, s+": "+b)
		}
	}
	return bad
}
