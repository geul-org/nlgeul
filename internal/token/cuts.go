//ff:func feature=token type=reader control=sequence
//ff:what Cuts — the word ends of a sentence read without a word table
package token

func Cuts(sentence string) []string {
	r, _ := New(Options{})
	return r.Cuts(sentence)
}
