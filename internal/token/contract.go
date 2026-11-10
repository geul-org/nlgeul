//ff:type feature=token type=model
//ff:what contract — a fused form whose first word is fixed: the form, the word it stands for, and the particle under it (text, name, where its entry is found)
package token

// key and form find the particle's entry in the index.
type contract struct {
	form string
	stem string
	text string
	name string
	key  rune
}
