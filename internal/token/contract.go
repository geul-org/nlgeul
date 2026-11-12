//ff:type feature=token type=model
//ff:what contract — a fused form whose first word is fixed: the form, the word it stands for, the particle under it (text, name, where its entry is found), whether it is read only right after a modifying closing, and the letters its word is written with when they are a wrong way to write it ("" for a word written right)
package token

// key and form find the particle's entry in the index.
type contract struct {
	form  string
	stem  string
	text  string
	name  string
	key   rune
	adn   bool
	wrote string
}
