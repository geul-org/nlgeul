//ff:type feature=token type=model
//ff:what stemCandidate — one candidate of the stem of a word: the stem as it is printed and looked up, the kind, name and verdict it stands with (the name is empty and the verdict 0 for a new word), and the letters it is written with when it is written wrong
package token

type stemCandidate struct {
	kindCandidate
	stem  string
	wrote string
}
