//ff:func feature=token type=helper control=sequence
//ff:what corrected — a part's candidates printed with the form it was written in when that is a misspelt form of the part's word: the candidates, the sign and 오타정정: before the written form; the candidates as they are for a part written right
package token

// This is the one place that puts the mark of a misspelt form into a
// printed part.
func corrected(names string, written string) string {
	if written == "" {
		return names
	}
	return names + typoSign + typoName + ":" + written
}
