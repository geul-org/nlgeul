//ff:func feature=token type=helper control=iteration dimension=2
//ff:what jsonWordAt — the first word of a sentence set down as data whose letters are the text given, and whether there is one
package token

func jsonWordAt(d JSONSentence, text string) (JSONWord, bool) {
	for _, sg := range d.Segments {
		for _, w := range sg.Words {
			if w.Text == text {
				return w, true
			}
		}
	}
	return JSONWord{}, false
}
