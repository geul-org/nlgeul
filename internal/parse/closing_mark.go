//ff:func feature=parse type=reader control=sequence
//ff:what closingMark — a whole-syllable closing as printed: its pieces, with the joining 으 as a piece before them when it was read
package parse

// text is what was read in the word (으 included).
func closingMark(e ending, text string) string {
	pieces, ok := closingPieces[e.form]
	if !ok {
		pieces = e.form
	}
	if text != e.form {
		return "+으+" + pieces
	}
	return "+" + pieces
}
