//ff:func feature=token type=reader control=sequence
//ff:what Lines — the lines of a sentence read without a word table
package token

func Lines(sentence string, trace bool) []string {
	r, _ := New(Options{})
	return r.Lines(sentence, trace)
}
