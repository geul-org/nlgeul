//ff:func feature=token type=helper control=iteration dimension=2
//ff:what jsonWordsFit — whether every word of a sentence set down as data has its letters at its place among the letters of the sentence and at least one candidate
package token

func jsonWordsFit(d JSONSentence, letters []rune) bool {
	for _, seg := range d.Segments {
		for _, w := range seg.Words {
			if w.Text != string(letters[w.Span[0]:w.Span[1]]) || len(w.Candidates) == 0 {
				return false
			}
		}
	}
	return true
}
